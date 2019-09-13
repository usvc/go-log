# Log

Everything related to logging for Go applications

# Usage

## Creating a logger

```go
import (
  "gitlab.com/usvc/modules/go/log/pkg/logger"
)

var log = logger.New()
```

## Creating a logger that streams to a remote FluentD 

```go
import (
  "gitlab.com/usvc/modules/go/log/pkg/logger"
  fluenthook "gitlab.com/usvc/modules/go/log/pkg/hooks/fluentd"
)

var log = logger.New()
fluentHook := fluenthook.NewHook(&fluenthook.HookConfig{
  Host:                    "localhost",
  Port:                    24224,
  InitializeRetryCount:    10,
  InitializeRetryInterval: time.Second * 1,
  Levels: []logrus.Level{
    logrus.TraceLevel,
    logrus.DebugLevel,
    logrus.InfoLevel,
    logrus.WarnLevel,
    logrus.ErrorLevel,
    logrus.PanicLevel,
  },
  Tag: "tag",
}
log.AddHook(fluentHook)
```


# Runbook

This section is a **work-in-progress**.

## Running the example application

```sh
make run
```

## Building the example binary

```sh
make build
```

## Testing it manually

Start FluentD:

```sh
make fluent
```

Run the example application:

```sh
make run
```

# License

This project is licensed under the [MIT license](./LICENSE).
