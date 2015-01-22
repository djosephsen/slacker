package handlers

import (
	sl "github.com/djosephsen/slacker/slackerlib"
	"fmt"
	"time"
)

var ChannelID = sl.MessageHandler{
	Name: `Chores/Channel ID`,
	Usage:	`<botname> (what channel)|(ID *channel)|(channel *ID): prints the Slack ID of the current channel`,
	Method:  `RESPOND`,
	Pattern: `(?i)(what channel)|(ID *channel)|(channel *ID)`,
	Run: func(e *sl.Event, match []string){
		reply := fmt.Sprintf("Current channel ID is: %s",e.Channel)
		e.Reply(reply)
	},
}

var ListChores = sl.MessageHandler{
	Name:	`Chores/ListChores`,
	Usage:	`botname (list chores)|(chore list): lists all registered chores`,
	Method:  `RESPOND`,
	Pattern: `(?i)(list chores)|(chore list)`,
	Run: func(e *sl.Event, match []string){
		var reply string
		if len(e.Sbot.Chores) == 0{
			reply=`No chores have been registered (sorry?)`
		}else{
			reply=`Name  :small_blue_diamond:  Schedule  :small_blue_diamond:  Firing in  :small_blue_diamond: Current State`
			for _,c := range e.Sbot.Chores{
				reply = fmt.Sprintf("%s\n%s:small_blue_diamond:%s:small_blue_diamond:%v:small_blue_diamond:%s",reply,c.Name, c.Sched, c.Next.Sub(time.Now()), c.State)
			}
		}
		e.Reply(reply)
	},
}

var ManageChores = sl.MessageHandler{
	Name:	`Chores/ManageChores`,
	Usage: `botname (start|stop) chore [chorename]: stops or starts the named chore`,
	Method:  `RESPOND`,
	Pattern: `(?i)(start|stop) chore (.*)`,
	Run: func(e *sl.Event, match []string){
		var reply string
		act:=match[1]
		cname:=match[2]
		c:=sl.GetChoreByName(cname,e.Sbot)
		if c == nil{ 
			reply = fmt.Sprintf("Chore not found: %s",cname)
		}else{
			if act==`stop`{
				sl.StopChore(c)
			}else{
				c.Start(e.Sbot)
			}
			reply = fmt.Sprintf("%s\n%s:small_blue_diamond:%s:small_blue_diamond:%v:small_blue_diamond:%s",reply,c.Name, c.Sched, c.Next.Sub(time.Now()), c.State)
		}
		e.Reply(reply)
	},
}
