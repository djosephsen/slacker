package inithooks

import (
	"encoding/json"
	sl "github.com/djosephsen/slacker/slackerlib"
)

var CbTest = sl.StartupHook{
	Name:  `cbTest`,
	Usage: `test`,
	Run: func(bot *sl.Sbot) {
		sl.Logger.Info(`setting up callback for type: pong`)
		myChan := make(chan map[string]interface{})

		bot.Register(sl.Callback{
			Name:    `test callback`,
			Key:     `type`,
			Pattern: `pong`,
			Channel: myChan,
		})
		thingy := <-myChan
		jthingy, _ := json.Marshal(thingy)
		sl.Logger.Debug(`CALLBACK FIRED:: `, string(jthingy))
	},
}
