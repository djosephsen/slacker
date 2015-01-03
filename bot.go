package main

import (
	"github.com/gorilla/websocket"
	"os"
	"time"
	"encoding/json"
	"fmt"
)

//the top level instantiation of a slackerbot
type Bot struct{
	Name					string
	Ws     				*websocket.Conn
	MID	 				int32
	Config 				*Config
	ReadThread 			*ReadThread
	WriteThread 		*WriteThread
	Broker 				*Broker
	StartupHooks		*[]StartupHook
	ShutdownHooks		*[]ShutdownHook
	sigChan				chan os.Signal
}

type ReadThread struct{
	Bot					*Bot
	Chan					chan Event
}

func (r *ReadThread) Start(b *Bot){
	r.Bot=b
	e := Event{}
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
	Chan				chan Event
	OutputFilters		*[]OutputFilter
}

func (w *WriteThread) Start(b *Bot){
	w.Bot=b
	for {
		e := <-w.Chan
		e.ID = b.NextMID()
		if ejson, _ := json.Marshal(e); len(ejson) >= 16000 {
				e = Event{e.ID, e.Type, e.Channel, fmt.Sprintf("ERROR! Response too large. %v Bytes!", len(ejson)), "", "", b}
			}
			b.Ws.WriteJSON(e)
			time.Sleep(time.Second * 1)
		}
}

//TODO: Make this threadsafe
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
			*b.StartupHooks=append(*b.StartupHooks, thing.(StartupHook))
		case ShutdownHook:
			*b.ShutdownHooks=append(*b.ShutdownHooks, thing.(ShutdownHook))
		case OutputFilter:
			*b.WriteThread.OutputFilters=append(*b.WriteThread.OutputFilters, thing.(OutputFilter))
		default:
			Logger.Error(`sorry I don't know what a `,t, `is`)
		}
	}
}
