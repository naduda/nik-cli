package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

func InitLogger(logname string) (*log.Logger, error) {
	e, err := os.OpenFile(logname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	logger := log.New(e, "", log.Ldate|log.Ltime)
	logger.SetOutput(&lumberjack.Logger{
		Filename:   logname,
		MaxSize:    10, // megabytes after which new file is created
		MaxBackups: 5,  // number of backups
		MaxAge:     1,  // days
	})
	return logger, nil
}
