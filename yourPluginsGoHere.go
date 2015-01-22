package main

import(
	sl "github.com/djosephsen/slacker/slackerlib"
	"github.com/djosephsen/slacker/inithooks"
	"github.com/djosephsen/slacker/handlers"
	"github.com/djosephsen/slacker/chores"
)

func initPlugins(b *sl.Sbot) error{
	b.Register(inithooks.Hai)
	b.Register(inithooks.Bai)
	b.Register(handlers.Syn)
	b.Register(handlers.Help)
	b.Register(handlers.Braintest)
	b.Register(handlers.ListChores)
	b.Register(handlers.ManageChores)
	b.Register(handlers.ChannelID)
	b.Register(chores.RTMPing)
	return nil
}
