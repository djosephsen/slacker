package slackerlib

import (
	"fmt"
	"net/http"
)

func (bot *Sbot) StartHttp() {
	http.HandleFunc("/", httpHi)
	err := http.ListenAndServe(":"+bot.Config.Port, nil)
	if err != nil {
		Logger.Error(err)
	}
}

func httpHi(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "sup. I'm a slackerbot")
}
