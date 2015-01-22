package slackerlib

import (
	`regexp`
	`fmt`
	`strings`
)

type Broker struct{
	Sbot	 *Sbot
   PreFilters        []*InputFilter
   MessageHandlers   []*MessageHandler
   EventHandlers     []*GenericEventHandler
}

func (broker *Broker) Start(bot *Sbot){
	broker.Sbot = bot	
	Logger.Debug(`Broker Started`)
	for {
      select {
      	case event := <-bot.ReadThread.Chan:
				event.Sbot = bot
         	go broker.This(&event)
      }
   }
}

func (b *Broker) This(e *Event){
   //run the pre-handeler filters
	if b.PreFilters != nil{ 
   	for _,filter := range b.PreFilters{ //run the pre-handler filters
     		e=filter.Run(e)
   	}
	}
   switch e.Type {
      case `message`:
         go b.HandleMessage(e)
      default :
         go b.HandleEvent(e)
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

func (b *Broker) HandleEvent(e *Event){
	Logger.Debug(`Broker:: caught event, type: `, e.Type, ` text:`, e.Text)
	if b.EventHandlers == nil{ return }
	for _,handler := range b.EventHandlers{
		go handler.Run(e)
	}
}

type InputFilter struct {
	Name		string
	Usage		string
	Run		func(e *Event) *Event
}

type MessageHandler struct {
	Name		string
	Method	string
	Pattern	string
	Usage		string
	Run		func(e *Event, match []string)
}

type GenericEventHandler struct {
	Name		string
	Usage		string
	Run		func(e *Event)
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
