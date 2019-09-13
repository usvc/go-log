package fluentd

import (
	"fmt"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"

	constants "gitlab.com/usvc/modules/go/log/pkg/constants"
)

// clearQueue clears the current log entry queue by sending all of them at
// once. to run upon successful connection to fluentd
func clearQueue(hook *Hook) {
	defer hook.log.Trace("ended")
	hook.log.Trace("called")
	for i := 0; i < len(hook.queue); i++ {
		hook.log.Infof("sending queued message [%v/%v]", i+1, len(hook.queue))
		hook.send(formatEntry(hook.queue[i]))
	}
	hook.queue = []*logrus.Entry{}
}

// formatEntry outputs a raw data object that can be sent to hook.send
func formatEntry(entry *logrus.Entry) map[string]interface{} {
	data := map[string]interface{}{
		constants.FieldLevel:     entry.Level.String(),
		constants.FieldMessage:   entry.Message,
		constants.FieldTimestamp: entry.Time.UTC().Format(time.RFC822Z),
	}
	for key, value := range entry.Data {
		data[key] = value
	}
	if entry.Caller != nil {
		data[constants.FieldFile] = fmt.Sprintf("%s:%v", entry.Caller.File, entry.Caller.Line)
		data[constants.FieldFunction] = entry.Caller.Function
	}
	return data
}

// handleInitializationError defines the retry logic for initialize()
func handleInitializationError(err error, hook *Hook) {
	if hook.config.InitializeRetryCount > hook.retryCount {
		hook.log.Warnf("unable to initialize fluentd logger: '%s'", err)
		hook.log.Debugf("attempting again in %v...", hook.config.InitializeRetryInterval)
		<-time.After(hook.config.InitializeRetryInterval)
		initialize(hook)
	} else {
		hook.log.Errorf("failed to initialize fluentd logger after %v attempts: '%s'", hook.config.InitializeRetryCount, err)
		hook.log.Infof("following %v log entries could not be sent to fluentd:", len(hook.queue))
		for i := 0; i < len(hook.queue); i++ {
			hook.log.Infof("%v", formatEntry(hook.queue[i]))
		}
		panic(err)
	}
}

// initialize is a controller method that creates a fluent logger
// instance and populates the .instance property if a successful
// connection to a fluentd service is established
func initialize(hook *Hook) {
	hook.log.Trace("called")
	defer func() {
		if r := recover(); r != nil {
			handleInitializationError(r.(error), hook)
		}
		hook.log.Trace("ended")
	}()
	hook.log.Debugf("attempting to initialize (%v attempts left)", hook.config.InitializeRetryCount-hook.retryCount)
	hook.retryCount++
	hook.isInitialising = true
	var err error
	hook.log.Debugf("connecting to %s:%v...", hook.config.Host, hook.config.Port)
	if hook.instance, err = fluent.New(createFluentConfig(hook.config)); err != nil {
		hook.instance = nil
		panic(err)
	}
	hook.log.Debug("fluentd initialized successfully")
	hook.send(map[string]interface{}{
		"level":   "debug",
		"message": "fluentd initialized successfully",
	})
	hook.isInitialising = false
	if len(hook.queue) > 0 {
		clearQueue(hook)
	}
}
