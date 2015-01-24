# Calling into the Slack API Proper

Slacker provides a function called *MakeApiReq()* which is a convienent
abstraction for calling into and getting responses from the Slack API. Just
create a new(ApiRequest) and pass it to MakeApiReq(), and you'll get back an
ApiResponse from SlackHQ. Here's the ApiRequest def: 

```
type ApiRequest struct{
   URL      string
   Values   url.Values
   Bot      *Sbot
}
```

Lets see how it works in practice. Imagine for a moment that you work someplace
terrible, and have been tasked with installing a PCI-Compliance filter that
finds and deletes messages containing strings that appear to be credit card
numbers from your corporate slack team.

In the time between when you give notice and actually leave, you could
implement that as a message handler like so:

```
var ZOMGPCI = sl.MessageHandler{
   Name: `ZOMG PCI`,
   Usage: `Listens for credit card numbers; deletes messages containing them`,
   Method:  `HEAR`,
   Pattern: ccPAT,
   Run: delCCFunc,
}
```
... where ccPAT is a regex string that matches all the various types of credit cards like...

```
var ccPAT := `^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|6(?:011|5[0-9]{2})[0-9]{12}|(?:2131|1800|35\d{3})\d{11})$` // good luck ever fixing that suckers
```
and delCCFunc is a function that calls the slackAPI to delete any message that
triggers the handler:

```
func delCCFunc (e *sl.Event, match []string){
	//create a url.Values
	values:=new(url.Values)
	values.Set(`channel`, e.Channel)
	values.Set(`ts`, e.Ts)

	// create an api request
	var req = ApiRequest{
      URL: 		`https://slack.com/api/chat.delete`,
      Values: 	values,
      Bot: 		e.Sbot,
   }

	// send the request to slack
   resp, err := MakeAPIReq(req)

	// see what we got back
	var reply string
	if err != nil {
  		reply = fmt.Sprintf(`CC# alert! (but I couldn't delete it because of an HTTP error: %s)`,err)
	}else if ! resp.Ok{
  		reply = fmt.Sprintf(`CC# alert! (but I couldn't delete it because the Slack API returned an error: %s)`,resp.Error)
	}else{
		reply = `please be more careful about pasting credit card numbers into the chat rooms!`
	}
   // tell the channel what we did
	e.Reply(reply)
}
```

You can read more about [url.Values
here](http://golang.org/pkg/net/url/#Values). It's basically just a
*map[string][]string* with some helpful functions.  Consult the [SlackAPI
docs](https://api.slack.com/methods/chat.delete) for a up-to-date list of
methods and their corrisponding URL's. 
