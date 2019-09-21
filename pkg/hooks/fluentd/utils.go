package fluentd

import (
	"fmt"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"

	"github.com/usvc/go-log/pkg/constants"
)

// clearQueue clears the current log entry queue by sending all of them at
// once. to run upon successful connection to fluentd
func clearQueue(hook *Hook) {
	defer hook.trace("ended")
	hook.trace("called")
	for i := 0; i < len(hook.queue); i++ {
		hook.debugf("sending queued message [%v/%v]", i+1, len(hook.queue))
		if err := hook.send(formatEntry(hook.queue[i])); err != nil {
			hook.warnf("failed to send queued message '%v'", hook.queue[i])
		} else {
			hook.queue = spliceLogEntry(hook.queue, uint(i))
		}
	}
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
	if hasNoRetryLimit(hook) || !hasReachedRetryLimit(hook) {
		hook.warnf("unable to initialize fluentd logger (attempt %v/%v): '%s'", hook.retryCount, hook.config.InitializeRetryCount, err)
		hook.debugf("attempting again in %v...", hook.config.InitializeRetryInterval)
		<-time.After(hook.config.InitializeRetryInterval)
		initialize(hook)
	} else {
		hook.errorf("failed to initialize fluentd logger after %v attempts: '%s'", hook.config.InitializeRetryCount, err)
		hook.warnf("following %v log entries could not be sent to fluentd: [\n%v\n]", len(hook.queue), hook.queue)
		panic(err)
	}
}

// hasReachedRetryLimit returns true if there are still retries left
func hasReachedRetryLimit(hook *Hook) bool {
	return hook.retryCount > hook.config.InitializeRetryCount
}

// hasNoRetryLimit returns true if the HookConfig.InitializeRetryCount property
// is less than 0, which indicates that we should never stop trying to
// reconnect to a FluentD service
func hasNoRetryLimit(hook *Hook) bool {
	return hook.config.InitializeRetryCount < 0
}

// initialize is a controller method that creates a fluent logger
// instance and populates the .instance property if a successful
// connection to a fluentd service is established
func initialize(hook *Hook) {
	hook.trace("called")
	defer func() {
		if r := recover(); r != nil {
			handleInitializationError(r.(error), hook)
		}
		hook.trace("ended")
	}()
	hook.debugf("attempting to initialize (attempt: %v/%v)", hook.retryCount, hook.config.InitializeRetryCount)
	hook.retryCount++
	hook.isInitialising = true
	var err error
	hook.debugf("connecting to %s:%v...", hook.config.Host, hook.config.Port)
	if hook.instance, err = fluent.New(createFluentConfig(hook.config)); err != nil {
		hook.instance = nil
		panic(err)
	}
	hook.debugf("fluentd successfully initialized")
	hook.send(map[string]interface{}{
		constants.FieldLevel:   "debug",
		constants.FieldMessage: "fluentd initialized successfully",
	})
	hook.isInitialising = false
	if len(hook.queue) > 0 {
		clearQueue(hook)
	}
}

func spliceLogEntry(entries []*logrus.Entry, index uint) []*logrus.Entry {
	return append(entries[:index], entries[index+1:]...)
}
