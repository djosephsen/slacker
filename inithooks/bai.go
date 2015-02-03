package inithooks

import (
	sl "github.com/djosephsen/slacker/slackerlib"
)

var Bai = sl.ShutdownHook{
	Name: `Bai`,
	Usage:`log a friendly goodbye message on a graceful shutdown`,
	Run:	func(bot *sl.Sbot){
		bot.Say(`Welp, I just got SigTermd, peaceOUT! :fist:`)
		sl.Logger.Info(`Caught SigTerm, slacker shutting down...ZOMGBAI!!`)
	},
}
