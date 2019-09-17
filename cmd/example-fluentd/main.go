package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/usvc/go-log/pkg/constants"
	fluenthook "github.com/usvc/go-log/pkg/hooks/fluentd"
	"github.com/usvc/go-log/pkg/logger"
)

var log *logrus.Logger

func init() {
	log = logger.New()
	fluentLogger := logger.New()
	fluentLogger.SetLevel(logrus.TraceLevel)
	fluentHook := fluenthook.NewHook(&fluenthook.HookConfig{
		Host:                    constants.DefaultFluentDHost,
		Port:                    constants.DefaultFluentDPort,
		InitializeRetryCount:    10,
		InitializeRetryInterval: time.Second * 1,
		Levels:                  constants.DefaultHookLevels,
		Tag:                     "tag",
	}, fluentLogger)
	log.AddHook(fluentHook)
}

func main() {
	<-time.After(2 * time.Second)
	log.WithFields(map[string]interface{}{
		"hello": "world",
	}).Info("hello world")
	<-time.After(2 * time.Second)
}
