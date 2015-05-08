package handlers

import (
	sl "github.com/djosephsen/slacker/slackerlib"
	"math/rand"
	"time"
)

var Syn = sl.MessageHandler{
	Name:    `Syn`,
	Usage:   `listen for '<botname> (syn|ping)', respond accordingly`,
	Method:  `RESPOND`,
	Pattern: `(?i)(syn|ping)`,
	Run: func(e *sl.Event, match []string) {
		now := time.Now()
		rand.Seed(int64(now.Unix()))
		replies := []string{
			"yeah um.. pong?",
			"WHAT?! jeeze.",
			"what? oh, um SYNACKSYN? ENOSPEAKTCP.",
			"RST (lulz)",
			"64 bytes from go.away.your.annoying icmp_seq=0 ttl=42 time=42.596 ms",
			"hmm?",
			"ack. what?",
			"pong. what?",
			"yup. still here.",
			"super busy just now.. Can I get back to you in like 5min?",
		}
		e.Reply(replies[rand.Intn(len(replies)-1)])
	},
}
