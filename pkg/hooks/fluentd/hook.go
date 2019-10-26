package fluentd

import (
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"
	constants "github.com/usvc/go-log/pkg/constants"
)

// Hook implements the logrus.Hook interface
type Hook struct {
	// config contains the configuration for the hook
	config *HookConfig

	// instance points to the official fluent logger and is
	// populated by initialize()
	instance *fluent.Fluent

	// isInitialising is used internally between when initialize()
	// is called and when the instance has achieved a connection to
	// the fluentd service
	isInitialising bool

	// log is an additional logger to pass in that the hook
	// will use to send status updates if it cannot connect to a
	// fluentd service
	log Logger

	// queue is used to store log entries that reach the hook before
	// fluentd is ready or when the connection to a fluentd service
	// is lost
	queue []*logrus.Entry

	retryCount int
}

type IHook interface {
	Levels() []logrus.Level
	Fire(*logrus.Entry) error
	Close()
	getQueuedEntryAt(uint) *logrus.Entry
	getQueueLength() uint
	removeLogFromQueue(uint)
	shouldRetry() bool
	send(map[string]interface{}) error
	post(string, map[string]interface{})
	trace(...interface{})
	debugf(string, ...interface{})
	warnf(string, ...interface{})
	errorf(string, ...interface{})
}

// Levels implements the logrus.Hook interface
func (hook *Hook) Levels() []logrus.Level {
	defer hook.trace("ended")
	hook.trace("called")
	return hook.config.Levels
}

// Fire implements the logrus.Hook interface
func (hook *Hook) Fire(entry *logrus.Entry) error {
	defer hook.trace("ended")
	hook.trace("called")
	if hook.instance == nil {
		hook.queue = append(hook.queue, entry)
		if !hook.isInitialising {
			go initialize(hook)
		}
		return nil
	}
	hook.send(createLogFromEntry(entry))
	return nil
}

// Close closes the hook's fluentd instance
func (hook *Hook) Close() {
	defer hook.trace("ended")
	hook.trace("called")
	hook.instance.Close()
}

// getQueuedEntryAt retrieves the log entry at the index :index
func (hook *Hook) getQueuedEntryAt(index uint) *logrus.Entry {
	if uint(len(hook.queue)) <= index {
		return nil
	}
	return hook.queue[index]
}

// getQueueLength returns the length of the queue
func (hook *Hook) getQueueLength() uint {
	return uint(len(hook.queue))
}

// removeLogFromQueue removes the log entry at index :index
// from the queue
func (hook *Hook) removeLogFromQueue(index uint) {
	spliceLogEntry(hook.queue, index)
}

// shouldRetry returns true if there are still retries left
func (hook *Hook) shouldRetry() bool {
	return hook.config.InitializeRetryCount < 0 || 
		hook.retryCount <= hook.config.InitializeRetryCount
}

// send formats and posts the data to the remote fluentd instance asynchronously
func (hook *Hook) send(data map[string]interface{}) error {
	defer hook.trace("ended")
	hook.trace("called")
	level := "log"
	if levelProperty, ok := data[constants.FieldLevel].(string); ok {
		level = levelProperty
	}
	go hook.post(level, data)
	return nil
}

// post sends the data to the remote fluentd instance
func (hook *Hook) post(level string, data map[string]interface{}) {
	defer hook.trace("ended")
	hook.trace("called")
	if err := hook.instance.Post(level, data); err != nil {
		hook.warnf("log entry <%v> could not be sent: '%s'", data, err)
	}
}

// trace sends a trace levelled log with the specified :message
// if a logger exists, otherwise nothing is done
func (hook *Hook) trace(message ...interface{}) {
	if hook.log != nil {
		hook.log.Trace(message...)
	}
}

// debugf sends a debug levelled formatted log with the specified
// :message if a logger exists, otherwise nothing is done
func (hook *Hook) debugf(format string, others ...interface{}) {
	if hook.log != nil {
		hook.log.Debugf(format, others...)
	}
}

// warnf sends a warn levelled formatted log with the specified
// :message if a logger exists, otherwise nothing is done
func (hook *Hook) warnf(format string, others ...interface{}) {
	if hook.log != nil {
		hook.log.Warnf(format, others...)
	}
}

// errorf sends a error levelled formatted log with the specified
// :message if a logger exists, otherwise nothing is done
func (hook *Hook) errorf(format string, others ...interface{}) {
	if hook.log != nil {
		hook.log.Errorf(format, others...)
	}
}
