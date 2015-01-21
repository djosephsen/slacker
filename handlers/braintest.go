package handlers

import (
	sl "github.com/djosephsen/slacker/slackerlib"
	"fmt"
	"regexp"
)

var Braintest = sl.MessageHandler{
	Name: `Brain-Test`,
	Usage:`<botname> brain [set|get] <key> <value>`,
	Method: `RESPOND`,
	Pattern: `(?i:brain) ((?i)set|get) (\w+) *(\w*)$`,
	Run:	func(e *sl.Event, match []string){
		brain := *e.Sbot.Brain
		if err := brain.Open(); err!=nil{
			sl.Logger.Debug(err)
			return
		}
		defer brain.Close()
		cmd := match[1]
		key := match[2]
		if matched,_ := regexp.MatchString(`(?i)set`, cmd); matched{
			val:=match[3]
			if err := brain.Set(key, []byte(val)); err != nil{
				sl.Logger.Debug(err)
				return
			}
			ret, _:=brain.Get(key)
			reply:=fmt.Sprintf("OK!, *%s* set to *%s*", key, string(ret))
			e.Reply(reply)	
		}else{
			val, err := brain.Get(key)
			if err != nil{
				reply:=fmt.Sprintf("Sorry, something went wrong: %s", err)
			   e.Reply(reply)	
				sl.Logger.Debug(err)
				return
			}
			e.Reply(string(val))	
		}
	},
}
