package main

import (
	"github.com/gorilla/websocket"
	"time"
)

//the top level instantiation of a slackerbot
type Bot struct{
	MID	 *MID
	Ws     *websocket.Conn
	Config Config
	Chan	 *[]BotChannels 
	MessageHandlers		 *[]MessageHandler
	StartupHooks	 *[]StartupHook
	ShutdownHooks	 *[]ShutdownHook
	GenericInputEventHandlers	 *[]GenericInputEventHandler
	OutputHooks		 *[]OutputHook
}

type MID int //per the slack api these need to be unique per connection
func (m *MID) Next() MID{
	m += 1
	return *m //return a copy so the MID doesn't change after assingment to a MSG{}
}

func (b *bot) Register(things ...interface{}) error {
	for t := range things{
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
