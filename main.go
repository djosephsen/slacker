package main

import (
"fmt"
)

func main(){
	
	bot:=new(Bot)
	bot.MID = 0
	bot.Conf = newConfig()
	bot.Logger = newLogger()

  if bot.Ws,err:=getMeASocket(bot.Conf.Token); err != nil{
      bot.Logger.Error(err)
	}
	
	if err = initHooks(bot); err !=nil{
      bot.Logger.Error(err)
	}

	//run startup hooks
	for h:=range bot.

	//in/out loop

}
