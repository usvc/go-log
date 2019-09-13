package main

import (
	"time"

	"github.com/sirupsen/logrus"

	fluenthook "gitlab.com/usvc/modules/go/log/pkg/hooks/fluentd"
	"gitlab.com/usvc/modules/go/log/pkg/logger"
)

var log *logrus.Logger

func init() {
	log = logger.New()
	fluentLogger := logger.New()
	fluentLogger.SetLevel(logrus.TraceLevel)
	fluentHook := fluenthook.NewHook(&fluenthook.HookConfig{
		Host:                    "localhost",
		Port:                    24224,
		InitializeRetryCount:    10,
		InitializeRetryInterval: time.Second * 1,
		Levels: []logrus.Level{
			logrus.TraceLevel,
			logrus.DebugLevel,
			logrus.InfoLevel,
			logrus.WarnLevel,
			logrus.ErrorLevel,
			logrus.PanicLevel,
		},
		Tag: "tag",
	}, fluentLogger)
	log.AddHook(fluentHook)
}

var id = 0

func main() {
	done := make(chan bool, 1)
	go func(tick <-chan time.Time) {
		for {
			<-tick
			id++
			log.WithFields(map[string]interface{}{
				"hi": "world",
			}).Printf("id: %v", id)
		}
	}(time.Tick(2 * time.Second))
	<-done
}
