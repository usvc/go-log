package main

import (
	"time"

	"github.com/sirupsen/logrus"

	"gitlab.com/usvc/modules/go/log/pkg/constants"
	fluenthook "gitlab.com/usvc/modules/go/log/pkg/hooks/fluentd"
	"gitlab.com/usvc/modules/go/log/pkg/logger"
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
