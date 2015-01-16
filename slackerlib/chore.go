package slackerlib

import (
   "fmt"
   "time"
   "github.com/gorhill/cronexpr"
)

type Chore struct {
   Name     string
   Usage    string
   Sched    string
   Run      func(b *Bot)
   State    string
   Next     time.Time
   Timer    *time.Timer
}

//an abstraction for the benifit of time.AfterFunc's second argument
// (probably just me being dense)
type ChoreTrigger struct {
	Chore		*Chore
	Bot		*Bot
}

func (t *ChoreTrigger) Pull(){
   Logger.Debug("Triggered: ",t.Chore.Name)
   t.Chore.State="running"
   go t.Chore.Run(t.Bot)
   t.Chore.Start(t.Bot)
}

// Schedule the chores
func (bot *Bot) StartChores() error{
   for _, c := range *bot.Chores {
      c.Start(bot)
      Logger.Debug("Started chore: ",c.Name)
   }
   return nil
}

func (c *Chore) Start(bot *Bot) error{
   Logger.Debug("Re-Starting: ",c.Name)
   expr := cronexpr.MustParse(c.Sched)
   if expr.Next(time.Now()).IsZero(){
      Logger.Debug("invalid schedule",c.Sched)
      c.State=fmt.Sprintf("NOT Scheduled (invalid Schedule: %s)",c.Sched)
   }else{
      Logger.Debug("valid Schedule: ",c.Sched)
      c.Next = expr.Next(time.Now())
      dur := c.Next.Sub(time.Now())
         if dur>0{
            Logger.Debug("valid duration: ",dur)
            if c.Timer == nil{
               Logger.Debug("creating a new timer")
					trigger:=&ChoreTrigger{
						Chore: c,
						Bot:   bot,
					}
               c.Timer = time.AfterFunc(dur, trigger.Pull) // auto go-routine'd
            }else{
               Logger.Debug("pre-existing timer found, resetting to: ",dur)
               c.Timer.Reset(dur) // auto go-routine'd
            }
         c.State=fmt.Sprintf("Scheduled: %s",c.Next.String())
         }else{
            Logger.Debug("invalid duration",dur)
            c.State=fmt.Sprintf("Halted. (invalid duration: %s)",dur)
         }
      }
   Logger.Debug("all set! Chore: ",c.Name, "scheduled at: ",c.Next)
   return nil
}

func GetChoreByName(name string, bot *Bot) *Chore{
   for _, c := range *bot.Chores {
      if c.Name == name{
         return &c
      }else{
         Logger.Debug("chore not found: ",name)
      }
   }
   return nil
}

func (c *Chore)Kill() error{
   Logger.Debug(`Stopping: `,c.Name)
   c.State=`Halted by request`
   c.Timer.Stop()
   return nil
}
