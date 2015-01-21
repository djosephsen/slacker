package slackerlib

import (
   "github.com/ccding/go-logging/logging"
   "github.com/danryan/env"
   "os"
   "time"
)

// Config struct
type Config struct {
   Name        string `env:"key=SLACKER_NAME default=slackerbot"`
   LogLevel    string `env:"key=SLACKER_LOG_LEVEL default=info"`
   Token 	   string `env:"key=SLACKER_TOKEN default=info"`
   RedisURL 	string `env:"key=SLACKER_REDIS_URL"`
}

func newConfig() *Config {
   c := &Config{}
   env.MustProcess(c)
   return c
}

func newLogger() *logging.Logger {
   format := "%25s [%s] %8s: %s\n time,name,levelname,message"
   timeFormat := time.RFC3339
   level := logging.GetLevelValue(`INFO`)
   logger, _ := logging.WriterLogger("slacker", level, format, timeFormat, os.Stdout, true)
   return logger
}
