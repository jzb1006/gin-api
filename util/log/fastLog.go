package log

import (
	"go-gin-api/init/log"
)

func Info(test string)  {
	log.Logger.Info(test)
}

func Debug(test string)  {
	log.Logger.Debug(test)
}