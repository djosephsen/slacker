package main

import(
	sl "github.com/djosephsen/slacker/slackerlib"
	"github.com/djosephsen/slacker/inithooks"
	"github.com/djosephsen/slacker/handlers"
	"github.com/djosephsen/slacker/chores"
)

func initHooks(b *sl.Bot) error{
	b.Register(inithooks.Hai)
	b.Register(inithooks.Bai)
	b.Register(handlers.Syn)
	b.Register(chores.Ping)
	return nil
}
