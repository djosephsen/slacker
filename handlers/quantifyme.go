package handlers

import (
	"fmt"
	sl "github.com/djosephsen/slacker/slackerlib"
	"math/rand"
	"regexp"
	"time"
)

var QuantifyMe = sl.MessageHandler{
	Name:    `Quantify Me`,
	Usage:   `<botname> quantify <noun>: replies with a randomly generated quantification of funny but probably NSFW personal attributes like 'mads', 'fucks', horribleness and passive-aggresivity`,
	Method:  `RESPOND`,
	Pattern: `(?i)quantify (\S+)`,
	Run: func(e *sl.Event, match []string) {
		var reply string
		now := time.Now()
		rand.Seed(int64(now.Unix()))
		user := match[1]
		if isMe, _ := regexp.MatchString(`(?i)me`, user); isMe {
			user = `you`
		}
		states := []string{`passive aggressive`, `mads`, `fucks`, `horrible`}
		state := states[rand.Intn(len(states)-1)]
		switch state {
		case `horrible`, `passive aggressive`:
			if user == `you` {
				reply = fmt.Sprintf(`%s are currently %%%d.%04d %s`, user, rand.Intn(int(100)), rand.Intn(int(1000)), state)
			} else {
				reply = fmt.Sprintf(`%s is currently %%%d.%04d %s`, user, rand.Intn(int(100)), rand.Intn(int(1000)), state)
			}
		case `mads`:
			if user == `you` {
				reply = fmt.Sprintf(`%s are %d.%04d %s`, user, rand.Intn(int(4)), rand.Intn(int(1000)), state)
			} else {
				reply = fmt.Sprintf(`%s is %d.%04d %s`, user, rand.Intn(int(4)), rand.Intn(int(1000)), state)
			}
		case `fucks`:
			if user == `you` {
				reply = fmt.Sprintf(`%s give precisely %f %s`, user, rand.Float64(), state)
			} else {
				reply = fmt.Sprintf(`%s gives precisely %f %s`, user, rand.Float64(), state)
			}
		}

		e.Respond(reply)
	},
}
