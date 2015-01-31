# There's not much to configure

Slacker configures via environment variables like a good 12-factor app. The
available vars are: 

* SLACKER_NAME
* SLACKER_TOKEN
* SLACKER_LOG_LEVEL
* SLACKER_REDIS_URL
* PORT

### Only two are required: 

* *SLACKER_NAME*:  The name of your bot. (this should match whatever you set your bot's name to on the SlackHQ integrations page.
* *SLACKER_TOKEN*: Your bot's SlackHQ API Token. 

### A couple others are optional: 

* *SLACKER_REDIS_URL*: if you set this to a valid redis system, Slacker will use it as its [brain](brain.md)
* SLACKER_LOG_LEVEL*: You can set this to 'DEBUG' to get debug output, otherwise it'll default to 'INFO'

### And one you can ignore completely:

* *PORT*: If this is set, slacker will start a completely useless http server.
 This behavior exists solely to prevent Heroku from prematurely killing
 Slacker. Even when this is used, you don't have to set it. Ignore away.
