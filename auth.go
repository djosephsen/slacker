package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"encoding/json"
	"strings"
)

// Given an API token, get me a web socket from slack pleasthnx
func getMeASocket(token string) (*websocket.Conn, error) {
	rtmURL:=`https://slack.com/api/rtm.start"`
	resp, err := http.PostForm(rtmURL, url.Values{"token": {token}}) 
	if err != nil{
		return new(websocket.Conn), fmt.Errorf("no dice with rtm.start: %v", err)
	}

	defer resp.Body.Close()
	authResp:=new(AuthResponse)
	dec:=json.NewDecoder(resp.Body)
	err=dec.Decode(authResp)
	if err != nil {
		return new(websocket.Conn), fmt.Errorf("Couldn't decode json. ERR: %v", err)
	}
	wsURL := strings.Split(authResp.URL, "/")
	wsURL[2] = wsURL[2] + ":443"
	authResp.URL = strings.Join(wsURL, "/")

	var Dialer websocket.Dialer
	header := make(http.Header)
	header.Add("Origin", "http://localhost/")

	ws, resp, err := Dialer.Dial(authResp.URL, header) 
	if err != nil{
		return new(websocket.Conn), fmt.Errorf("no dice dialing that websocket: %v", err)
	}
	return ws, nil
}
