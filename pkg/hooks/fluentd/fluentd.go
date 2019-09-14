package fluentd

import (
	"os"

	"github.com/sirupsen/logrus"
)

// log is the basic default logger to be used if no logger is provided
var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: &logrus.TextFormatter{},
	Level:     logrus.WarnLevel,
}

// NewHook instantiates a new minimal fluentd hook given a configuration
// and an optional logger to use (if a logger is not provided, the in-
// built one will be used)
func NewHook(
	config *HookConfig,
	logger ...*logrus.Logger,
) *Hook {

	if len(logger) > 0 {
		log = logger[0]
	}
	return &Hook{
		config: config,
		log:    log,
		queue:  []*logrus.Entry{},
	}
}
