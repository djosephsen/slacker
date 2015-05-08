package chores

import (
	sl "github.com/djosephsen/slacker/slackerlib"
)

var RTMPing = sl.Chore{
	Name:  `rtm-ping`,
	Usage: `Sends an RTM ping message to slack every 20 seconds`,
	Sched: `*/20 * * * * * *`,
	Run: func(bot *sl.Sbot) {
		bot.Send(&sl.Event{
			ID:   0, //the write thread will give this a real MID
			Type: `ping`,
			Text: `just pingin`,
		})
	},
}
