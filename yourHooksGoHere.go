package main

import(
	sl "github.com/djosephsen/slacker/slackerlib"
	"github.com/djosephsen/slacker/sdhooks"
)

func initHooks(b *sl.Bot) error{
	//b.Register(chores.Ping)
	b.Register(sdhooks.Bai)
	b.Register(sdhooks.Hai)
	return nil
}
