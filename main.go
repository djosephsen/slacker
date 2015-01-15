package main

import (
sl "github.com/djosephsen/slacker/slackerlib"
"os/signal"
"syscall"
)

func main(){
	
	bot := new(sl.Bot)
	err := bot.Init()

	go bot.WriteThread.Start(bot)
	go bot.ReadThread.Start(bot)
	go bot.Broker.Start(bot)

	//initialize the handlers, chores and filters
	if err = initHooks(bot); err !=nil{
      sl.Logger.Error(err)
	}

	//run startup-hooks
	if bot.StartupHooks != nil{
		for _,h := range *bot.StartupHooks{
			go h.Run(bot)
		}
	}

	signal.Notify(bot.SigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
  	stop := false
   for !stop {
      select {
      case sig := <-bot.SigChan:
         switch sig {
         case syscall.SIGINT, syscall.SIGTERM:
            stop = true
         }
      }
   }
   // Stop listening for new signals
   signal.Stop(bot.SigChan)

	// shutdownn hooks
	if bot.ShutdownHooks != nil{
		for _,h := range *bot.ShutdownHooks{
			h.Run(bot)
		}
	}
}
