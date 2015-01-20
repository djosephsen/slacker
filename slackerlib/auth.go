package slackerlib

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"encoding/json"
	"strings"
	"io/ioutil"
)

// Go forth and get a websocket and all the Slack Team Metadata 
func (bot *Sbot) getMeASocket() error {
	rtmURL:=`https://slack.com/api/rtm.start`
	resp, err := http.PostForm(rtmURL, url.Values{"token": {bot.Config.Token}}) 
	if err != nil{
		return fmt.Errorf("no dice with rtm.start: %v", err)
	}
	defer resp.Body.Close()

	authResp:=new(AuthResponse)
	dec:=json.NewDecoder(resp.Body)
	err=dec.Decode(authResp)
	if err != nil {
		out,_ := ioutil.ReadAll(resp.Body)
		fmt.Printf(`%s`,string(out))
		return fmt.Errorf("Couldn't decode json. ERR: %v", err)
	}

	if authResp.URL == ""{
		out,_ := ioutil.ReadAll(resp.Body)
		fmt.Printf(`Auth failure: %s`,string(out))
		return fmt.Errorf("Auth failure: %s",string(out))
	}

	wsURL := strings.Split(authResp.URL, "/")
	wsURL[2] = wsURL[2] + ":443"
	authResp.URL = strings.Join(wsURL, "/")
	Logger.Debug(`Team Wesocket URL: `, authResp.URL)

	var Dialer websocket.Dialer
	header := make(http.Header)
	header.Add("Origin", "http://localhost/")

	ws, resp, err := Dialer.Dial(authResp.URL, header)
	if err != nil{
		return fmt.Errorf("no dice dialing that websocket: %v", err)
	}

	//yay we're websocketing
	bot.Ws=ws
	bot.Meta=authResp
	return nil
}
