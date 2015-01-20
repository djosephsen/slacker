#Slacker Plugins

Typically, chatbots are expected to respond to commands and conversation
happening in the chatroom. The usual arrangement is that you create a plug-in,
that specifies a regular expression for the bot to trigger on, and how you want
the bot to respond.  Because Slacker is a chatops bot, there are a few
different kinds of plugins you can define. The classical listen/respond
behavior, for example, is implemented in Slacker by a [Message
Handler](handlers.md) plugin. 

Here are the four classes of plugins you can currently configure: 

 * [Handlers](handlers.md) respond to events recieved from Slack's API. Events include messages (people saying things to each other) as well as other things like notifications (psst so and so is typing things), pongs (responses to pings), and well, events (psst so and so has uploaded a cat gif).

 * [Chores](chres.md) allow you to tell the bot to do things on a periodic schedule. If you wanted Slacker to, for example grab the RSS feed from lobste.rs every hour and parse it for new articles, you could implement a Chore to do this.

 * [Filters](filters.md) can rewrite events as they come in, or go out to Slack. If it was Snoop doggy-dog's birthday, and you wanted everything Slacker said on that day to be translated into [gangsta](http://www.gizoogle.net/textilizer.php) you could implement an OuputFilter.

 * [InitHooks](hooks.md) are serially executed whenever Slacker starts up or shuts down. If you need your chatops bot to do some initialization work whenever it starts up or shuts down, you can configure the appropriate inithook for it. 

Lets walk through the creation and addition of a simple plugin. Say you wanted
your chatbot to respond "MMMMMMMM omgbacon" whenever it heard the word "bacon"
in conversation. The first step would be to cd to the handlers directory, and
create a file called bacon.go that looked like this: 

'''
package handlers

import (
   sl "github.com/djosephsen/slacker/slackerlib"
)

func Bacon

var Syn = sl.MessageHandler{
   Name: `Bacon`,
   Usage:`listen for 'bacon', respond with 'MMMMMMM omgbacon'`,
   Method: `HEAR`,
   Pattern: `bacon`,
   Run:  func(e *sl.Event, match []string){
      e.Reply(`ack`)
   },
}
'''
