package mlog

import (
	"github.com/orandin/lumberjackrus"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

var (
	Logger *logrus.Logger
)

func InitLogrus(path, infoName, errorName, traceName string, level uint32, caller bool) {
	Logger = logrus.New()
	Logger.SetReportCaller(caller) // true 打印日志位置,默认是false
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetLevel(logrus.Level(level))
	Logger.AddHook(newRotateHook(path, infoName, errorName, traceName))
}

func newRotateHook(path, infoName, errorName, traceName string) logrus.Hook {
	hook, err := lumberjackrus.NewHook(
		&lumberjackrus.LogFile{
			Filename:   path + infoName,
			MaxSize:    20,
			MaxAge:     10,
			MaxBackups: 10,
			LocalTime:  true,
			Compress:   true,
		},
		logrus.InfoLevel,
		&logrus.TextFormatter{DisableColors: true, TimestampFormat: time.DateTime},
		&lumberjackrus.LogFileOpts{
			logrus.TraceLevel: &lumberjackrus.LogFile{
				Filename: path + traceName,
			},
			logrus.ErrorLevel: &lumberjackrus.LogFile{
				Filename:   path + errorName,
				MaxSize:    20,
				MaxBackups: 10,
				MaxAge:     10,
				Compress:   true,
				LocalTime:  true,
			},
		},
	)
	if err != nil {
		log.Fatalln("logrus 日志插件初始化失败！", err)
	}
	return hook
}
