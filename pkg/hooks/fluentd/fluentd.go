package fluentd

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: &logrus.TextFormatter{},
	Level:     logrus.DebugLevel,
}

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

func NewHookConfig(
	host string,
	port int,
) *HookConfig {
	return &HookConfig{
		Host: host,
		Port: port,
	}
}
