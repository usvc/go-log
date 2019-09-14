/*
Package fluentd exposes a class that adheres to the logrus.Hook interface. Create it,
then add it to an existing logger using the .AddHook method of a logrus.Logger

Example initialization:

	```go
	import (
		"gitlab.com/usvc/modules/go/log/pkg/logger"
		"gitlab.com/usvc/modules/go/log/pkg/hooks/fluentd"
		"gitlab.com/usvc/modules/go/log/pkg/constants"
	)

	// ...

	logger := logger.New()
	fluentHook := fluentd.NewHook(&fluenthook.HookConfig{
		Host:                    "localhost",
		Port:                    24224,
		InitializeRetryCount:    10,
		InitializeRetryInterval: time.Second * 1,
		Levels:                  constants.DefaultHookLevels,
		Tag:                     "tag",
	}
	logger.AddHook(fluentHook)
	```

For the full guide visit https://gitlab.com/usvc/modules/go/log
*/
package fluentd
