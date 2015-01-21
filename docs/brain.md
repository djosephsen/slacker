#Sbot.Brain

Slacker automatically creates and provides an interface to a persistent storage
back-end (aka a *Brain*) every time it starts up. If you do nothing at all,
your bot's Brain is backed by an in-memory data structure that will go away
every time your bot is rebooted.

That's ok for development work, and for plugins that don't need to remember
things between reboots, but it's obviously of limited utility. To enable a
Redis-Backed Brain, you need only specify your redis URL as an environment
variable: 

`export SLACKER_REDIS_URL='redis://passwd@192.168.0.1:6379'`

## Usage
Sbot.Brain is an interface with the following definition: 

```
type Brain interface {
   Open() error
   Close() error
   Get(string) ([]byte, error)
   Set(key string, data []byte) error
   Delete(string) error
}
```
generally speaking, the pattern is for your plugin to:
	1. call sBot.Brain.Open() to connect to the back-end
   2. Set, Get, or Delete values to its hearts content
	3. call sBot.Brain.Close() to close the connection to the back-end when its done

Checkout the [braintest](/handlers/braintest.go) handler for a more detailed example. 
