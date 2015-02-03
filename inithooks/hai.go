package inithooks

import (
	sl "github.com/djosephsen/slacker/slackerlib"
)

var Hai = sl.StartupHook{
	Name: `Hai`,
	Usage:`log a friendly hello message on startup`,
	Run:	func(bot *sl.Sbot){
		sl.Logger.Info(`Oh Hai!.  We're all initialized and ready to go!`)
		bot.Say(`Oh Hai! I just woke up! :hand:`)
	},
}
