package main

import (
sl "github.com/djosephsen/slacker/slackerlib"
"os/signal"
"syscall"
)

func main(){
	
	//make a bot
	bot := new(sl.Sbot)
	err := bot.Init()

	//start the read, write and broker threads
	go bot.WriteThread.Start(bot)
	go bot.ReadThread.Start(bot)
	go bot.Broker.Start(bot)

	//Read in and register all the handlers, chores, filters, and hooks
	if err = initPlugins(bot); err !=nil{
      sl.Logger.Error(err)
		return
	}

	//run startup-hooks
	if bot.StartupHooks != nil{
		for _,h := range bot.StartupHooks{
			go h.Run(bot)
		}
	}

	// Start the chores
	if bot.Chores != nil{
		for _,c := range bot.Chores{
			c.Start(bot)
		}
	}

	// Loop
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

	// Run Shutdownn Hooks
	if bot.ShutdownHooks != nil{
		for _,h := range bot.ShutdownHooks{
			h.Run(bot)
		}
	}

	//wait for the write thread to stop (so the shutdown hooks have a chance to run)
	bot.WriteThread.RunChan <- true
	<- bot.SyncChan
}
