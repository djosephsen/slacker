package slackerlib

import (
	"github.com/gorilla/websocket"
	"github.com/ccding/go-logging/logging"
	"os"
	"time"
	"encoding/json"
	"fmt"
	"strings"
)

var Logger = newLogger()

//the top level instantiation of a slackerbot
type Sbot struct{
	Name					string
	Ws     				*websocket.Conn
	MID	 				int32
	Config 				*Config
	Meta					*AuthResponse
	ReadThread 			*ReadThread
	WriteThread 		*WriteThread
	Broker 				*Broker
	Brain					*Brain
	StartupHooks		[]*StartupHook
	ShutdownHooks		[]*ShutdownHook
	Chores				[]*Chore
	SigChan				chan os.Signal
	SyncChan				chan bool
}

func (bot *Sbot) Init() error {
	var err error
	bot.MID = 0
	bot.Config = newConfig()
	bot.Name = bot.Config.Name
	Logger.SetLevel(logging.GetLevelValue(strings.ToUpper(bot.Config.LogLevel)))
	bot.SigChan = make(chan os.Signal, 1)
	bot.SyncChan = make(chan bool)
	bot.WriteThread = &WriteThread{
		Chan:		make(chan Event),
		RunChan:	make(chan bool),
	}
	bot.ReadThread = &ReadThread{
		Chan:		make(chan Event,1),
	}
	bot.Broker = &Broker{
		Sbot:					bot,
	}
	bot.Brain, err = bot.NewBrain()
	if err != nil{
		return err
	}
	err = bot.getMeASocket()
	if err != nil{
		return err
	}
	Logger.Debug(`Joined team: `, bot.Meta.Team.Name )
	return nil
}

type ReadThread struct{
	Sbot					*Sbot
	Chan					chan Event
}

func (r *ReadThread) Start(b *Sbot){
	r.Sbot = b
	e := Event{}
	Logger.Debug(`Read-Thread Started`)
	for {
		b.Ws.ReadJSON(&e)
		if (e != Event{}) { // if the event isn't empty
			e.Sbot = b
			b.ReadThread.Chan <- e
			e = Event{}
		}
	}
}

type WriteThread struct{
	Sbot				*Sbot
	OutputFilters	[]*OutputFilter
	Chan				chan Event
	RunChan			chan bool
}

func (w *WriteThread) Start(b *Sbot){
	w.Sbot=b
	Logger.Debug(`Write-Thread Started`)
	stop := false
	for !stop {
		select{
		case e := <-w.Chan:
			e.Sbot=nil //nil the bot pointer out or Marshal() dies horrible infinite recusive death
			e.ID = b.NextMID()
			Logger.Debug(`WriteThread:: Outbound `,e.Type,` channel: `,e.Channel,`. text: `,e.Text)
			if ejson, _ := json.Marshal(e); len(ejson) >= 16000 {
				e = Event{
				ID: e.ID, 
				Type: e.Type, 
				Channel: e.Channel, 
				Text: fmt.Sprintf("ERROR! Response too large. %v Bytes!", len(ejson)), 
				}
			}
				b.Ws.WriteJSON(e)
				time.Sleep(time.Second * 1)
		case stop = <- w.RunChan:
			stop = true
			}
		}
	b.SyncChan <- true
}

//probably need to make this thread-safe (for now only the write thread uses it)
func (b *Sbot) NextMID() int32{
	b.MID += 1
	Logger.Debug(`incrementing MID to `, b.MID)
	return b.MID
}

func (b *Sbot) Register(things ...interface{}){
	for _,thing := range things{
		switch t := thing.(type) {
		case MessageHandler:
			m:=thing.(MessageHandler)
			Logger.Debug(`registered MessageHandler: `,m.Name)
			b.Broker.MessageHandlers=append(b.Broker.MessageHandlers, &m)	
		case GenericEventHandler:
			g:=thing.(GenericEventHandler)
			Logger.Debug(`registered Event Handler: `,g.Name)
			b.Broker.EventHandlers=append(b.Broker.EventHandlers, &g)
		case InputFilter:
			i:=thing.(InputFilter)
			Logger.Debug(`registered Input Filter: `, i.Name)
			b.Broker.PreFilters=append(b.Broker.PreFilters, &i)
		case StartupHook:
			s:=thing.(StartupHook)
			Logger.Debug(`registered StartupHook: `, s.Name)
			b.StartupHooks=append(b.StartupHooks, &s)
		case ShutdownHook:
			s:=thing.(ShutdownHook)
			Logger.Debug(`registered ShutdownHook: `, s.Name)
			b.ShutdownHooks=append(b.ShutdownHooks, &s)
		case OutputFilter:
			o:=thing.(OutputFilter)
			Logger.Debug(`registered OutputFilter: `, o.Name)
			b.WriteThread.OutputFilters=append(b.WriteThread.OutputFilters, &o)
		case Chore:
			c:=thing.(Chore)
			Logger.Debug(`registered Chore: `,c.Name)
			b.Chores=append(b.Chores, &c)
		default:
			weirdType:=fmt.Sprintf(`%T`,t)
			Logger.Error(`sorry I cant register this handler because I don't know what a `,weirdType, ` is`)
		}
	}
}

// Say something in the named channel (or the default channel if none specified)
func (b *Sbot) Say(s string, channel ...string){
	var c string
	if channel != nil{
		c=channel[0]
	}else{
		c=b.DefaultChannel()
	}
	event := Event{
		Type: 	`message`,
		Channel: c,
		Text:		s,
		}
	b.WriteThread.Chan <- event
}

//returns the Team's default channel 
func (b *Sbot) DefaultChannel() string{
	for _, c := range b.Meta.Channels{
		if c.IsGeneral{
			return c.ID
		}
	}
	return b.Meta.Channels[0].ID
}
