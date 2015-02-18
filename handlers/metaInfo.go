package handlers

import (
	sl "github.com/djosephsen/slacker/slackerlib"
	"fmt"
	"regexp"
)

var MetaInfo = sl.MessageHandler{
	Name: `MetaInfo`,
	Usage:	`<botname> (channel|user) (list|dump) <id>:: examine meta-info on channels and users`,
	Method:  `RESPOND`,
	Pattern: `(?i)(channel|user)s* (list|dump) *(\w+)*`,
	Run: func(e *sl.Event, match []string){
		sl.Logger.Debug(`metaInfo:: called`)
		typeOfThing:=match[1]
		cmd := match[2]
		id := match[3]
		var reply string
		if matches,_ := regexp.MatchString(`(?i)list`,cmd); matches{
			reply=listThing(e.Sbot,typeOfThing)
		}else if matches,_ := regexp.MatchString(`(?i)dump`,cmd); matches{
			reply=dumpThing(e.Sbot, typeOfThing, id)
		}
		if reply != ``{
			e.Reply(reply)
		}
	},
}

func listThing(bot *sl.Sbot,typeOfThing string) string{
	var reply string
	if matches,_ := regexp.MatchString(`(?i)channel`,typeOfThing); matches{
		reply=`Channels:`
		for _,c:=range bot.Meta.Channels{
			reply=fmt.Sprintf("%s\n%s (%s)",reply, c.ID, c.Name)
		}
	}else if matches,_ := regexp.MatchString(`(?i)user`,typeOfThing); matches{
		reply=`Users:`
		for _,u:=range bot.Meta.Users{
			reply=fmt.Sprintf("%s\n%s (%s)",reply, u.ID, u.Name)
		}
	}
	return reply
}

func dumpThing(bot *sl.Sbot,typeOfThing string, id string) string{
	var reply string
	if matches,_ := regexp.MatchString(`(?i)channel`,typeOfThing); matches{
		channel:=bot.Meta.GetChannel(id)
		reply:=fmt.Sprintf("Channel: %s", channel.Name)
		reply=fmt.Sprintf("%s\n%s",reply, channel)
	}else if matches,_ := regexp.MatchString(`(?i)user`,typeOfThing); matches{
		user:=bot.Meta.GetUser(id)
		reply:=fmt.Sprintf("User: %s", user.Name)
		reply=fmt.Sprintf("%s\n%s",reply, user)
	}
	return reply
}
