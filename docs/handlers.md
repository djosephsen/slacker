#Event Handlers

Events are pushed from SlackHQ to Slacker via a websocket connection as they
occur. Every event has a [*type*](https://api.slack.com/events). By far the
most common type of event recieved from Slack are 'message' events, which
represent people typing messages to each other.

At the moment, you can register two different types of Event handlers with
Slacker: a *MessageHandler*, which is an event handler specifically designed to
handle message-type events, and a *EventHandler* which handles, well,
everything else. 

The event handler definition looks like this: 

```
type MessageHandler struct {
   Name     string
   Method   string
   Pattern  string
   Usage    string
   Run      func(e *Event, match []string)
}
```

... whereas the EventHandler looks like this: 

```
type EventHandler struct {
   Name     string
   Usage    string
   Type     string
   Run      func(p *HandlerPackage)
}
```

Message handlers are easy. You register a regex pattern (as Pattern), and
slacker will send you events that match the pattern you specify, along with a
[]string of substrings that specifically matched your pattern.

If, for example you set Pattern to '(\w) (\w) (\w)', and someone in the
chatroom were to say 'foo bar biz', then the contents of the *match* slice
would look like this: 

```
match[0] == `foo bar biz`
match[1] == `foo`
match[2] == `bar`
match[3] == `biz`
```

The event struct contains the full event that occured with details like who
said whatever matched your pattern, and what channel they said it in. The event
struct also has a pointer back to the top-level bot struct, so you can do
pretty much anything you want. This is the event def: 

```
type Event struct {
   ID      int32    `json:"id,omitempty"`
   Type    string `json:"type,omitempty"`
   Channel string `json:"channel,omitempty"`
   Text    string `json:"text,omitempty"`
   User    string `json:"user,omitempty"`
   UserName    string `json:"username,omitempty"`
   BotID    string `json:"bot_id,omitempty"`
   Subtype  string `json:"subtype,omitempty"`
   Ts      string `json:"ts,omitempty,omitempty"`
   Sbot    *Sbot
}
```
