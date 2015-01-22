package handlers

import sl "github.com/djosephsen/slacker/slackerlib"

var Tableflip = sl.MessageHandler{
	Name: `Tableflip`,
	Usage: `bot flips a unicode table whenever it overhears '(table)*flip(table)*'`,
	Method:  `HEAR`,
	Pattern: `(?i)(table)*flip(table)*`,
	Run: func(e *sl.Event, match []string) {
		e.Respond(`(╯°□°）╯︵ ┻━┻`)
	},
}
