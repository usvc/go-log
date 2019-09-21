package fluentd

import (
	"os"

	"github.com/sirupsen/logrus"
)

// createDefaultLogger creates the basic default logger to be used if
// no logger is provided
func createDefaultLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.TextFormatter{},
		Level:     logrus.WarnLevel,
	}
}

// NewHook instantiates a new minimal fluentd hook given a configuration
// and an optional logger to use (if a logger is not provided, the in-
// built one will be used)
func NewHook(
	config *HookConfig,
	logger ...Logger,
) *Hook {
	var hookLogger Logger
	if len(logger) > 0 {
		hookLogger = logger[0]
	} else {
		hookLogger = createDefaultLogger()
	}
	return &Hook{
		config: config,
		log:    hookLogger,
		queue:  []*logrus.Entry{},
	}
}
