package fluentd

import (
	"fmt"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"

	"github.com/usvc/go-log/pkg/constants"
)

// addCustomDataToLog creates the data field in the passed in :log
func addCustomDataToLog(log map[string]interface{}, entry *logrus.Entry) {
	entryData := map[string]interface{}{}
	for key, value := range entry.Data {
		entryData[key] = value
	}
	log[constants.FieldData] = entryData
}

// addCallerDataToLog creates the file and function field in the passed in :log
func addCallerDataToLog(log map[string]interface{}, entry *logrus.Entry) {
	log[constants.FieldFile] = fmt.Sprintf("%s:%v", entry.Caller.File, entry.Caller.Line)
	log[constants.FieldFunction] = entry.Caller.Function
}

// clearQueue clears the current log entry queue by sending all of them at
// once. to run upon successful connection to fluentd
func clearQueue(hook IHook) {
	defer hook.trace("ended")
	hook.trace("called")
	queueLength := hook.getQueueLength()
	for i := uint(0); i < queueLength; i++ {
		hook.debugf("sending queued message [%v/%v]", i+1, queueLength)
		queuedEntry := hook.getQueuedEntryAt(i)
		logToSend := createLogFromEntry(queuedEntry)
		if err := hook.send(logToSend); err != nil {
			hook.warnf("failed to send queued message '%v'", logToSend)
		} else {
			hook.removeLogFromQueue(i)
		}
	}
}

// createBaseLogFromEntry creates a basic log entry that can be sent to hook.send
func createBaseLogFromEntry(entry *logrus.Entry) map[string]interface{} {
	return map[string]interface{}{
		constants.FieldLevel:     entry.Level.String(),
		constants.FieldMessage:   entry.Message,
		constants.FieldTimestamp: entry.Time.UTC().Format(constants.TimestampFormat),
	}
}

// createLogFromEntry outputs a raw data object that can be sent to hook.send
func createLogFromEntry(entry *logrus.Entry) map[string]interface{} {
	data := createBaseLogFromEntry((entry))
	if entry.Data != nil {
		addCustomDataToLog(data, entry)
	}
	if entry.Caller != nil {
		addCallerDataToLog(data, entry)
	}
	return data
}

// handleInitializationError defines the retry logic for initialize()
func handleInitializationError(err error, hook *Hook) {
	if hook.shouldRetry() {
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
	hook.send(createBaseLogFromEntry(&logrus.Entry{
		Level:   logrus.DebugLevel,
		Message: "fluentd initialized successfully",
		Time:    time.Now(),
	}))
	hook.isInitialising = false
	if len(hook.queue) > 0 {
		clearQueue(hook)
	}
}

// spliceLogEntry removes the index :index from the slice of log entries
// provided (:entries)
func spliceLogEntry(entries []*logrus.Entry, index uint) []*logrus.Entry {
	return append(entries[:index], entries[index+1:]...)
}
