package fluentd

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/usvc/go-log/pkg/constants"
	formatters "github.com/usvc/go-log/pkg/formatters/logrus"
)

// createDefaultLogger creates the basic default logger to be used if
// no logger is provided
func createDefaultLogger() *logrus.Logger {
	return &logrus.Logger{
		Formatter:    formatters.Text,
		Out:          os.Stderr,
		Level:        logrus.WarnLevel,
		ReportCaller: true,
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

// NewConfig returns a hook configuration with the default configuration
func NewConfig() *HookConfig {
	return &HookConfig{
		Host:                    constants.DefaultFluentDHost,
		Port:                    constants.DefaultFluentDPort,
		InitializeRetryCount:    constants.DefaultInitializeRetryCount,
		InitializeRetryInterval: constants.DefaultInitializeRetryInterval,
		Levels:                  constants.DefaultHookLevels,
		Tag:                     constants.DefaultFluentDTag,
	}
}
