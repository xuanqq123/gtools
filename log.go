package gtools

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
  )

func LogInit(logPath, logFile string) {
	if logPath == "" {
		logPath = "."
	}
	if logFile == "" {
		logFile = "run.log"
	}
  logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   logPath + "/" + logFile,
		MaxSize:    5,    
		MaxBackups: 10,   
		MaxAge:     3,  
		Compress:   true, 
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Error("Log message")
}
