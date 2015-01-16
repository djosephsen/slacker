package chores

import (
	sl "github.com/djosephsen/slacker/slackerlib"
)

var Ping = sl.Chore{
	Name:  `ping`,
	Usage:  `Sends an RTM ping message to slack every 20 seconds`,
	Sched: `*/20 * * * * * *`,
	Run: func(bot *sl.Bot){
		event:=sl.Event{
			ID: 0,  //the write thread will give this a real MID
			Type: `ping`,
			Text: `just pingin`,
		}
		bot.WriteThread.Chan <- event
	},
}
