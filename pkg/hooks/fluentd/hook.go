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
	hook.send(formatEntry(entry))
	return nil
}

// Close closes the hook's fluentd instance
func (hook *Hook) Close() {
	defer hook.trace("ended")
	hook.trace("called")
	hook.instance.Close()
}

// send posts the data to the remote fluentd instance
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
