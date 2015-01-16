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
type Bot struct{
	Name					string
	Ws     				*websocket.Conn
	MID	 				int32
	Config 				*Config
	Meta					*AuthResponse
	ReadThread 			*ReadThread
	WriteThread 		*WriteThread
	Broker 				*Broker
	StartupHooks		*[]StartupHook
	ShutdownHooks		*[]ShutdownHook
	SigChan				chan os.Signal
	SyncChan				chan bool
}

func (bot *Bot) Init() error {
   bot.MID = 0
   bot.Config = newConfig()
	bot.Name = bot.Config.Name
	bot.SigChan = make(chan os.Signal, 1)
	bot.SyncChan = make(chan bool)
	bot.WriteThread=&WriteThread{
		Chan:		make(chan Event),
		RunChan:	make(chan bool),
	}
	bot.ReadThread=&ReadThread{
		Chan:		make(chan Event,1),
	}
	bot.Broker=new(Broker)
	err:=bot.getMeASocket()
   if err != nil{
      return err
   }
	bot.StartupHooks=new([]StartupHook)
	bot.ShutdownHooks=new([]ShutdownHook)
	Logger.SetLevel(logging.GetLevelValue(strings.ToUpper(bot.Config.LogLevel)))
	Logger.Debug(`Joined team: `, bot.Meta.Team.Name )
	return nil
}

type ReadThread struct{
	Bot					*Bot
	Chan					chan Event
}

func (r *ReadThread) Start(b *Bot){
	r.Bot=b
	e := Event{}
	Logger.Debug(`Read-Thread Started`)
	for {
		b.Ws.ReadJSON(&e)
		if (e != Event{}) { // if the event isn't empty
			e.Bot = b
			b.ReadThread.Chan <- e
			e = Event{}
		}
	}
}

type WriteThread struct{
	Bot				*Bot
	OutputFilters	*[]OutputFilter
	Chan				chan Event
	RunChan			chan bool
}

func (w *WriteThread) Start(b *Bot){
	w.Bot=b
	w.OutputFilters = new([]OutputFilter)
	Logger.Debug(`Write-Thread Started`)
	stop := false
	for !stop {
		select{
		case e := <-w.Chan:
			e.Bot=nil //nil this out or Marshal() dies horrible infinite recusive death
			e.ID = b.NextMID()
			Logger.Debug(`WriteThread:: Outbound `,e.Type,`. text: `,e.Text)
			if ejson, _ := json.Marshal(e); len(ejson) >= 16000 {
				e = Event{e.ID, e.Type, e.Channel, fmt.Sprintf("ERROR! Response too large. %v Bytes!", len(ejson)), "", "", b}
			}
				b.Ws.WriteJSON(e)
				time.Sleep(time.Second * 1)
		case stop = <- w.RunChan:
			stop = true
			}
		}
	b.SyncChan <- true
}

//probably need to make this threadsafe
func (b *Bot) NextMID() int32{
	b.MID += 1
	Logger.Debug(`incrementing MID to `, b.MID)
	return b.MID
}

func (b *Bot) Register(things ...interface{}){
	for _,thing := range things{
		switch t := thing.(type) {
		case MessageHandler:
			*b.Broker.MessageHandlers=append(*b.Broker.MessageHandlers,thing.(MessageHandler))	
		case GenericEventHandler:
			*b.Broker.EventHandlers=append(*b.Broker.EventHandlers, thing.(GenericEventHandler))
		case PreHandlerFilter:
			*b.Broker.PreFilters=append(*b.Broker.PreFilters, thing.(PreHandlerFilter))
		case StartupHook:
			suHook:=thing.(StartupHook)
			Logger.Debug(`registered StartupHook: `,suHook.Name)
			*b.StartupHooks=append(*b.StartupHooks, thing.(StartupHook))
		case ShutdownHook:
			sdHook:=thing.(ShutdownHook)
			Logger.Debug(`registered ShutdownHook: `,sdHook.Name)
			*b.ShutdownHooks=append(*b.ShutdownHooks, sdHook)
		case OutputFilter:
			*b.WriteThread.OutputFilters=append(*b.WriteThread.OutputFilters, thing.(OutputFilter))
		default:
			weirdType:=fmt.Sprintf(`%T`,t)
			Logger.Error(`sorry I cant register this handler because I don't know what a `,weirdType, ` is`)
		}
	}
}

func (b *Bot) Say(s string){
	event := Event{
		Type: 	`message`,
		Channel: b.Meta.Channels[0].ID,
		Text:		s,
		}
	b.WriteThread.Chan <- event
}
