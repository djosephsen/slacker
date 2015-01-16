package handlers

import (
	sl "github.com/djosephsen/slacker/slackerlib"
)

var Syn = sl.MessageHandler{
	Name: `Syn`,
	Usage:`listen for '<botname> syn', respond with 'ack'`,
	Method: `RESPOND`,
	Pattern: `syn`,
	Run:	func(e *sl.Event, match []string){
		e.Reply(`ack`)	
	},
}
