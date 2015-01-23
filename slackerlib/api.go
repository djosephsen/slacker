package slackerlib

import(
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"encoding/json"
	"strings"
)


type ApiRequest struct{
	URL		string
	Values	url.Values
	Bot		*Sbot
}

//base function for communicating with the slack api
func makeAPIReq(req ApiRequest)(*ApiResponse, error){
	resp:=new(ApiResponse)
	req.Values.Set(`token`, req.Bot.Config.Token)

	reply, err := http.PostForm(req.URL, req.Values)
   if err != nil{
      return resp, err
   }
   defer reply.Body.Close()

	dec := json.NewDecoder(reply.Body)
   err = dec.Decode(resp)
	if err != nil {
		return resp, fmt.Errorf("Couldn't decode json. ERR: %v", err)
	}
	return resp, nil
}

// Go forth and get a websocket for RTM and all the Slack Team Metadata
func (bot *Sbot) getMeASocket() error {
   var req = ApiRequest{
      URL: `https://slack.com/api/rtm.start`,
		Values: make(url.Values),
      Bot: bot,
   }
   authResp,err := makeAPIReq(req)
   if err != nil{
      return err
   }

   if authResp.URL == ""{
      return fmt.Errorf("Auth failure")
   }
   wsURL := strings.Split(authResp.URL, "/")
   wsURL[2] = wsURL[2] + ":443"
   authResp.URL = strings.Join(wsURL, "/")
   Logger.Debug(`Team Wesocket URL: `, authResp.URL)

   var Dialer websocket.Dialer
   header := make(http.Header)
   header.Add("Origin", "http://localhost/")

   ws, _, err := Dialer.Dial(authResp.URL, header)
   if err != nil{
      return fmt.Errorf("no dice dialing that websocket: %v", err)
   }

   //yay we're websocketing
   bot.Ws=ws
   bot.Meta=authResp
   return nil
}


// parses sBot.Meta to return a user's Name field given its ID
func (meta *ApiResponse) GetUserName(id string) string{
   for _,user := range meta.Users{
      if user.ID == id{
         return user.Name
      }
   }
   return ``
}

// convinience function to reply to a message event
func (event *Event) Reply(s string){
   replyText:=fmt.Sprintf(`%s: %s`, event.Sbot.Meta.GetUserName(event.User), s)
   response := Event{
      Type:    event.Type,
      Channel: event.Channel,
      Text:    replyText,
      }
   event.Sbot.WriteThread.Chan <- response
}

// convinience function to respond to a message event
func (event *Event) Respond(s string){
   response := Event{
      Type:    event.Type,
      Channel: event.Channel,
      Text:    s,
      }
   event.Sbot.WriteThread.Chan <- response
}
