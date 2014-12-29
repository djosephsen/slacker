package main

import (
"fmt"
)

func main(){
	
	bot:=new(Bot)
	bot.MID = 0
	bot.Conf = newConfig()
	bot.Logger = newLogger()

  if bot.Ws, err := getMeASocket(bot.Conf.Token); err != nil{
      bot.Logger.Error(err)
	}
	
	//initialize the handlers, chores and filters
	if err = initHooks(bot); err !=nil{
      bot.Logger.Error(err)
	}

	//run startup-hooks
	for h:=range bot.

	go bot.ReadThread.Start(bot)
	go bot.WriteThread.Start(bot)

	// read loop
	for {
		select {
		case event := <-bot.ReadThread.Chan:
			go bot.BrokerEvent(event) //in handleMessage.go
		}
	}
}
