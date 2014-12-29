import (
   "fmt"
   "github.com/ccding/go-logging/logging"
   "github.com/danryan/env"
   "os"
   "strings"
   "time"
)

// Config struct
type config struct {
   Name        string `env:"key=SLACKER_NAME default=slackerbot"`
   StoreName   string `env:"key=SLACKER_BRAIN default=memory"`
   LogLevel    string `env:"key=SLACKER_LOG_LEVEL default=info"`
   Token 	   string `env:"key=SLACKER_TOKEN default=info"`
   PreFilters        *[]PreHandlerFilter
   MessageHandlers   *[]MessageHandler
   EventHandlers     *[]GenericEventHandler
   PostFilters       *[]PostHandlerFilter
}

func newConfig() *config {
   c := &config{}
   env.MustProcess(c)
   return c
}

func newLogger() *logging.Logger {
   format := "%25s [%s] %8s: %s\n time,name,levelname,message"
   timeFormat := time.RFC3339
   levelStr := strings.ToUpper(Config.LogLevel)
   level := logging.GetLevelValue(levelStr)
   logger, _ := logging.WriterLogger("slacker", level, format, timeFormat, os.Stdout, true)
   return logger
}

