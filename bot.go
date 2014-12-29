package main

import (
	"github.com/gorilla/websocket"
	"time"
)

//the top level instantiation of a slackerbot
type Bot struct{
	Name					string
	Ws     				*websocket.Conn
	Config 				Config
	ReadThread 			*ReadThread
	WriteThread 		*WriteThread
	PreFilters	 		*[]PreHandlerFilter
	MessageHandlers	*[]MessageHandler
	EventHandlers		*[]GenericEventHandler
	PostFilters	 		*[]PostHandlerFilter
}

type ReadThread struct{
	Chan					Channel
}

func (r *ReadThread) Start(b *bot){
	e := Event{}
	for {
		b.Ws.ReadJSON(&e)
		if (e != Event{}) { // if the event isn't empty
			b.ReadThread.Chan <- e
			e = Event{}
		}
	}
}

type WriteThread struct{
	MID	 			*MID
	OutFilters		*[]OutputFilter
	Chan				Channel
}

func (w *WriteThread) Start(b *bot){
	for {
		e := <-w.Chan
		if eJson, _ := json.Marshal(e); len(eJson) >= 16000 {
				ejson = Event{e.Id, e.Type, e.Channel, fmt.Sprintf("ERROR! Response too large. %v Bytes!", len(ejson)), "", ""}
			}
			b.Ws.WriteJSON(ejson)
			time.Sleep(time.Second * 1)
		}
}

type MID int //per the slack api these need to be unique per connection
func (m *MID) Next() MID{
	m += 1
	return *m //return a copy so the MID doesn't change after assingment to a MSG{}
}

func (b *bot) BrokerEvent(e *event) error {
	for filter := range b.PreFilters{ //run the pre-handler filters
		e=filter(e)
	}
	switch e.Type {
		case `message`:
		/
		default :
	}
}


func (b *bot) Register(things ...interface{}) error {
	for thing := range things{
		switch t := thing.(type) {
			case MessageHandler
				bot.MessageHandlers=append(thing, bot.MessageHandlers)	
			case GenericInputEventHandler
				bot.GenericInputEventHandlers=append(thing, bot.GenericInputEventHandlers)	
			case StartupHook
				bot.StartupHooks=append(thing, bot.StartupHooks)	
			case ShutdownHook
				bot.ShutdownHooks=append(thing, bot.ShutdownHooks)	
			case OutputHook
				bot.OutputHooks=append(thing, bot.OutputHooks)	
			default 
		}
	}
}
