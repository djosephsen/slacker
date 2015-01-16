package sdhooks

import (
	sl "github.com/djosephsen/slacker/slackerlib"
)

var Bai = sl.ShutdownHook{
	Name: `Bai`,
	Usage:`log a friendly goodbye message on a graceful shutdown`,
	Run:	func(b *sl.Bot){
		sl.Logger.Info(`Caught SigTerm, slacker shutting down...ZOMGBAI!!`)
	},
}
