package slackerlib

import (
	`regexp`
	`fmt`
	`strings`
	`encoding/json`
)

// broker reads messages from the slack websocket and brokers them out
type Broker struct{
	Sbot	 *Sbot
   PreFilters        []*InputFilter
   MessageHandlers   []*MessageHandler
   EventHandlers     []*EventHandler
}

func (broker *Broker) Start(bot *Sbot){
	broker.Sbot = bot	
	Logger.Debug(`Broker Started`)
	for {
		thingy := make(map[string]interface{})
		bot.Ws.ReadJSON(&thingy)
      go broker.This(thingy)
   }
}

//Figure out what kind of thingy this is, package it, and ship it to the right handler(s)
func (b *Broker) This(thingy map[string]interface{}){
//run the pre-handeler filters
	if b.PreFilters != nil{ 
   	for _,filter := range b.PreFilters{ //run the pre-handler filters
     		thingy = filter.Run(thingy)
   	}
	}
	// stop here if a prefilter delted our thingy
	if len(thingy) == 0 { return }

	jthingy,_ := json.Marshal(thingy)
	typeOfThingy := thingy[`type`]
	switch typeOfThingy{
	case `message`:
		message := new(Event)
		json.Unmarshal(jthingy, message)
		message.Sbot = b.Sbot
      b.HandleMessage(message)
	default:
		b.HandleWTF(thingy)
	}
}

func (b *Broker) HandleMessage(e *Event){
	Logger.Debug(`Broker:: caught message, text: `, e.Text)
	if b.MessageHandlers == nil{ return }
	botNamePat := fmt.Sprintf(`^(?:@?%s[:,]?)\s+(?:${1})`, e.Sbot.Name)
	for _,handler := range b.MessageHandlers{
		var r *regexp.Regexp
		if handler.Method == `RESPOND`{
			r = regexp.MustCompile(strings.Replace(botNamePat,"${1}", handler.Pattern, 1))
		}else{
			r = regexp.MustCompile(handler.Pattern)
		}
		if r.MatchString(e.Text){
			match:=r.FindAllStringSubmatch(e.Text, -1)[0]
		   Logger.Debug(`Broker:: running handler: `, handler.Name)
			go handler.Run(e, match) 
		}
	}
}

func (b *Broker) HandleWTF(thingy map[string]interface{}){
	Logger.Debug(`Broker:: caught unknown type: `,thingy[`type`])
	if b.EventHandlers == nil{ return }
	for _,handler := range b.EventHandlers{
		if handler.Type == `*`{
			handler.Run(&HandlerPackage{Type: `unk`,Sbot: b.Sbot, Thingy: thingy})
		}
	}
}

type InputFilter struct {
	Name		string
	Usage		string
	Run		func(thingy map[string]interface{}) map[string]interface{}
}

type MessageHandler struct {
	Name		string
	Method	string
	Pattern	string
	Usage		string
	Run		func(e *Event, match []string)
}

type EventHandler struct {
	Name		string
	Usage		string
	Type		string
	Run		func(pack *HandlerPackage)
}

type HandlerPackage struct {
	Type		string
	Sbot		*Sbot
	Thingy	interface{}
}

type OutputFilter struct {
	Name		string
	Usage		string
	Run		func(e *Event)
}

type StartupHook struct {
	Name		string
	Usage		string
	Run		func(b *Sbot)
}

type ShutdownHook struct {
	Name		string
	Usage		string
	Run		func(b *Sbot)
}
