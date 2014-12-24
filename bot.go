package main

import (
	"github.com/gorilla/websocket"
	"time"
)

//the top level instantiation of a slackerbot
type Bot struct{
	MID	 *MID
	Ws     *websocket.Conn
	Config *Config
	Chan	 *[]BotChannels 
}

type MID int //per the slack api these need to be unique per connection
func (m *MID) Next() MID{
	m += 1
	return *m //return a copy so the MID doesn't change after assingment to a MSG{}
}


