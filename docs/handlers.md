#Event Handlers

See the [plugins](plugins.md) page for info on getting started writing handlers
and other types of plugins for slacker. 

Events are pushed from SlackHQ to Slacker via a websocket connection as they
occur. Every event has a [*type*](https://api.slack.com/events). By far the
most common type of event recieved from Slack are 'message' events, which
represent people typing messages to each other.

At the moment, you can register two different types of Event handlers with
Slacker: a *MessageHandler*, which is an event handler specifically designed to
handle message-type events, and an *EventHandler* which handles, well,
everything else. 

The eventhandlers and message handlers act a little differently. The former
expect you to specify the type of event you want to handle (this is a regex, so
`.` matches every type. When the broker recieves a non-message type event, from
Slack, it matches the event's *type* string against your pattern, and if it
matches, your handler's Run() function is executed.  

Event handlers are passed a

The general idea is, you
