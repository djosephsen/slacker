package main

import (
"github.com/ccding/go-logging/logging"
"os/signal"
"strings"
"syscall"
)

var (
	Logger = newLogger()
)

func main(){
	
	bot:=new(Bot)
	bot.MID = 0
	bot.Config = newConfig()
	Logger.SetLevel(logging.GetLevelValue(strings.ToUpper(bot.Config.LogLevel)))

  var err error
  bot.Ws, err = getMeASocket(bot.Config.Token) 
  if err != nil{
      Logger.Error(err)
		return
	}
	
	go bot.WriteThread.Start(bot)
	go bot.ReadThread.Start(bot)
	go bot.Broker.Start(bot)

	//initialize the handlers, chores and filters
	if err = initHooks(bot); err !=nil{
      Logger.Error(err)
	}

	//run startup-hooks
	for _,h := range *bot.StartupHooks{
		go h.Run(bot)
	}

	signal.Notify(bot.sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
  	stop := false
   for !stop {
      select {
      case sig := <-bot.sigChan:
         switch sig {
         case syscall.SIGINT, syscall.SIGTERM:
            stop = true
         }
      }
   }
   // Stop listening for new signals
   signal.Stop(bot.sigChan)

	// shutdownn hooks
	for _,h := range *bot.ShutdownHooks{
		go h.Run(bot)
	}
}
