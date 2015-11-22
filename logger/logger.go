package logger

import (
	"github.com/ArthurHlt/gominlog"
	"log"
)

var logger *gominlog.MinLog

func GetLogger() *log.Logger {
	return logger.GetLogger()
}

func SetLogger(loggerLog *log.Logger) {
	logger.SetLogger(loggerLog)
}


func GetMinLog() *gominlog.MinLog {
	if logger == nil {
		logger = gominlog.NewClassicMinLogWithPackageName("microserv-helper")
	}
	return logger
}