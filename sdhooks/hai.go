package sdhooks

import (
	sl "github.com/djosephsen/slacker/slackerlib"
)

var Hai = sl.StartupHook{
	Name: `Hai`,
	Usage:`log a friendly hello message on startup`,
	Run:	func(b *sl.Bot){
		sl.Logger.Info(`Oh Hai!.  We're all initialized and ready to go!`)
	},
}
