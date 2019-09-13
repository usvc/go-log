package fluentd

import (
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/sirupsen/logrus"
)

type HookConfig struct {
	// Host contains the hostname of the fluentd service
	//
	// example: "fluentd.monitoring.svc.cluster.local"
	Host string

	// Port contains the port that the fluentd service
	// is listening on
	//
	// example: 24224
	Port int

	// InitializeRetryCount indicates how many times the
	// initialize() function should try to connect to the
	// fluentd service before it fails
	//
	// example: 10
	InitializeRetryCount int

	// InitializeRetryInterval indicates the duration in
	// between connection attempts to fluentd by initialize()
	//
	// example: 5 * time.Second
	InitializeRetryInterval time.Duration

	// Levels is an array of logrus levels that the hook
	// should be activated for
	//
	// example: []logrus.Level{logrus.TraceLevel}
	Levels []logrus.Level

	// Tag defines the base tag used to tag the log entries
	// sent to fluentd
	//
	// example: "application"
	Tag string
}

// createFluentConfig creates a configuration object for the official
// fluent-logger-golang package
//
// ref: https://github.com/fluent/fluent-logger-golang/blob/master/fluent/fluent.go
func createFluentConfig(config *HookConfig) fluent.Config {
	return fluent.Config{
		FluentHost:   config.Host,
		FluentPort:   config.Port,
		MaxRetry:     100,
		MaxRetryWait: 1000,
		RequestAck:   false,
		RetryWait:    1000,
		TagPrefix:    config.Tag,
		Timeout:      5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
}
