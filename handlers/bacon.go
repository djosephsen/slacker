package handlers

import (
	sl "github.com/djosephsen/slacker/slackerlib"
)

func myRunFunc (e *sl.Event, match []string){
	e.Respond(`MMMMMMMMmmm ... omgbacon`)
}

var Bacon = sl.MessageHandler{
	Name: `Bacon`,
	Usage:`listen for 'bacon', respond with 'MMMMMMmmmm... omgbacon'`,
	Method: `HEAR`,
	Pattern: `bacon`,
	Run:		myRunFunc,
}
