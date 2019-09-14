# Log

Everything related to logging for Go applications. The main logger this work is based on is [`sirupsen/logrus`](https://github.com/sirupsen/logrus) and packages in this repository provides a standardised, best-practices logger.

Contained in here are three main packages:

- Logger (`import "gitlab.com/usvc/modules/go/log/pkg/logger"`)
- FluentD Hook (`import "gitlab.com/usvc/modules/go/log/pkg/hooks/fluentd"`)
- Constants (`import "gitlab.com/usvc/modules/go/log/pkg/constants"`)

- - -

# Usage


## Creating a logger

```go
import (
  "gitlab.com/usvc/modules/go/log/pkg/logger"
)

var log = logger.New()
```


## Creating a logger that streams to a remote FluentD 

> For the following to work, you'll need a FluentD service reachable
> at `localhost:24224`

```go
import (
  "gitlab.com/usvc/modules/go/log/pkg/logger"
  "gitlab.com/usvc/modules/go/log/pkg/hooks/fluentd"
  "gitlab.com/usvc/modules/go/log/pkg/constants"
)

var log = logger.New()
fluentHook := fluentd.NewHook(&fluentd.HookConfig{
  Host:                    constants.DefaultFluentDHost,
  InitializeRetryCount:    10,
  InitializeRetryInterval: time.Second * 1,
  Levels:                  constants.DefaultHookLevels,
  Port:                    constants.DefaultfluentDPort,
  Tag: "tag",
}
log.AddHook(fluentHook)
```

- - -

# License

This project is licensed under the [MIT license](./LICENSE).

- - -

# Development Runbook

This section contains notes for working on code in this project/contributing.


## Directory structure

The `cmd` directory contains example usages compilable to actual binaries.

The `lib` directory contains things that are not intended for export or use by external code.

The `pkg` directory contains all the exported things.


## Starting FluentD

You will need a FluentD service exposed to your local environment at `localhost:24224`. To get it up, run:

```sh
make fluentd
```


## Running the example application

The following runs the example with the most features:

```sh
make run
```

For other example applications, see the [`Makefile`](./Makefile)


## Testing

To run the tests:

```sh
make test
```

- - -

# Cheers
