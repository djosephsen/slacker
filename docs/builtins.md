#Built-in Plugins

Slacker comes with a slew of built-in plugins of various types. You can see
what plugins are currently registered and running by asking slacker for help. 

![help](docs/screenshots/help.png)

The Built-in message handlers are pretty self explanitory. There are a few
included plugins that overhear and respond to conversation, and you may find
this annoying. These are the IKR plugin, which listens for keywords that convey
enthuasim and the flip table plugin, which flips a table at any mention of the
words flip and table in any order in any single sentence. 

![annoying](docs/screenshots/annoying.png)

You should know that a few of Slackers internal processes are implemented as
chores and event handlers. One example of this is the rtm-ping plugin which
implements the RTM pings [required by the upstream Slack API]. If you disable
this chore, Slacker will stop sending RTM pings to SlackHQ and may wind up
having keep-alive problems as a result. 
