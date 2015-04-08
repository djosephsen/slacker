package handlers

import (
	"time"
	"math/rand"
	"strings"
	sl "github.com/djosephsen/slacker/slackerlib"
)

var IKR = sl.MessageHandler{
	Name: `IKR (I know, RIGHT?!)`,
	Usage: `listens for enthusiasm; responds with validation`,
	Method: `HEAR`,
	Pattern: genPattern(),
	Run: ikrRun,
}

func genPattern() string{
	//build a HYOOGE messy regex of stuff we want to match on
   triggers:=[]string{
		"best.*ev(er|ar)",
		"so good",
		"they have the best",
		"awesome",
		"I love",
		"fantastic|wonderful|outstanding|magnificent|brilliant|genius|amazing|epic|nice!",
		"ZOMG|OMG|OMFG",
		"(so|pretty) great",
		"off the hook",
	}
	pat := "(?i)"+strings.Join(triggers,"|")
	return pat
}

func ikrRun(e *sl.Event, match []string) {
   now:=time.Now()
	rand.Seed(int64(now.Unix()))
	//be less annoying by only firing half the time
 	if rand.Intn(100) >= 50{
		return 
	}
		replies := []string{
			"*I know right?!*",
			"*OMG* couldn't agree more",
			":+1:",
			"+1",
			":arrow_up: THAT",
			":arrow_up: you complete me :arrow_up:",
			"so true",
			"agreed.",
			"that's the fact jack",
			"YUUUUUUP",
			"that's what I'm talkin bout",
			"*IKR?!*",
			"singit",
			"^droppin the truth bombs :boom: :boom: :boom:",
			"#legit",
			"_nods emphatically in agreement_",
			"for REALZ though",
			"FOR REALSIES",
			"it's like you *literally* just read my mind right now",
	}
	reply := replies[rand.Intn(len(replies)-1)]
	e.Respond(reply)
}
