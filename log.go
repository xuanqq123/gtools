package gtools

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LevelFileHook struct {
	Levs       []logrus.Level
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	Formatter  logrus.Formatter
	Logger     *lumberjack.Logger
}

func NewLevelFileHook(levels []logrus.Level, filename string, maxSize, maxBackups, maxAge int, compress bool, formatter logrus.Formatter) *LevelFileHook {
	return &LevelFileHook{
		Levs:       levels,
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
		Formatter:  formatter,
		Logger: &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
			Compress:   compress,
		},
	}
}

func (hook *LevelFileHook) Fire(entry *logrus.Entry) error {
	logString, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.Logger.Write(logString)
	return err
}

func (hook *LevelFileHook) Levels() []logrus.Level {
	return hook.Levs
}

func LogInit(logPath, logFile string) {
	if logPath == "" {
		logPath = "."
	}

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	infoHook := NewLevelFileHook(
		[]logrus.Level{logrus.InfoLevel},
		fmt.Sprintf("%s/%s_info.log", logPath, logFile),
		5, 10, 3, true,
		&logrus.JSONFormatter{},
	)

	errorHook := NewLevelFileHook(
		[]logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel},
		fmt.Sprintf("%s/%s_err.log", logPath, logFile),
		5, 10, 3, true,
		&logrus.JSONFormatter{},
	)

	logrus.AddHook(infoHook)
	logrus.AddHook(errorHook)
	logrus.SetLevel(logrus.DebugLevel)
}
