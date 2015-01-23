#Event Handlers

Events are pushed from SlackHQ to Slacker via a websocket connection as they
occur. Every event has a [*type*](https://api.slack.com/events). By far the
most common type of event recieved from Slack are 'message' events, which
represent people typing messages to each other.

At the moment, you can register two different types of Event handlers with
Slacker: a *MessageHandler*, which is an event handler specifically designed to
handle message-type events, and a *GenericEventHandler* which handles, well,
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

... whereas the GenericEventHandler looks like this: 

```
type GenericEventHandler struct {
   Name     string
   Usage    string
   Run      func(e *Event)
}
```
