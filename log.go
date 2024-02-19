package gtools

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
  )

func LogInit() {
  logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   LogPath + LogFile,
		MaxSize:    5,    
		MaxBackups: 10,   
		MaxAge:     3,  
		Compress:   true, 
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Error("Log message")
}
