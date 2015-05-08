package handlers

import (
	"encoding/json"
	sl "github.com/djosephsen/slacker/slackerlib"
)

var MetaUpdater = sl.EventHandler{
	Name:  `Meta Updater`,
	Usage: `keeps Sbot.Meta up to date using event traffic from Slackhq`,
	Type:  `.`,
	Run:   metaUpdaterFunc,
}

func metaUpdaterFunc(hp *sl.HandlerPackage) {
	sl.Logger.Debug(`MetaUpdater:: got a: `, hp.Type)
	//jthingy,_:=json.Marshal(hp.Thingy)
	thingy := hp.Thingy
	switch hp.Type {
	case `group_join`:
		newChannel(hp.Sbot, thingy[`channel`])
	case `channel_created`:
		newChannel(hp.Sbot, thingy[`channel`])
	case `team_join`:
		newUser(hp.Sbot, thingy[`user`].(sl.User))
	}
}

func newChannel(bot *sl.Sbot, chanThingy interface{}) {
	channel := new(sl.Channel)
	jthingy, _ := json.Marshal(chanThingy)
	json.Unmarshal(jthingy, channel)
	for _, exists := range bot.Meta.Channels {
		if exists.ID == channel.ID {
			sl.Logger.Debug(`pre-existing channel: `, channel.Name)
			return
		}
	}
	sl.Logger.Debug(`adding channel to Meta: `, channel.Name)
	bot.Meta.Channels = append(bot.Meta.Channels, *channel)
}

func newUser(bot *sl.Sbot, user sl.User) {
	for _, exists := range bot.Meta.Users {
		if exists.ID == user.ID {
			sl.Logger.Debug(`pre-existing user: `, user.Name)
			return
		}
	}
	sl.Logger.Debug(`adding user to Meta: `, user.Name)
	bot.Meta.Users = append(bot.Meta.Users, user)
}

/*type HandlerPackage struct {
   Type     string
   Sbot     *Sbot
   Thingy   map[string]interface{}
}
*/
