# Log

[![pipeline status](https://gitlab.com/usvc/modules/go/log/badges/master/pipeline.svg)](https://gitlab.com/usvc/modules/go/log/commits/master)

Everything related to logging for Go applications. The main logger this work is based on is [`sirupsen/logrus`](https://github.com/sirupsen/logrus) and packages in this repository provides a standardised, best-practices logger.

Contained in here are three main packages:

- Logger (`import "github.com/usvc/go-log/pkg/logger"`)
- FluentD Hook (`import "github.com/usvc/go-log/pkg/hooks/fluentd"`)
- Constants (`import "github.com/usvc/go-log/pkg/constants"`)

- - -

# Usage


## Creating a logger

```go
import (
  "github.com/usvc/go-log/pkg/logger"
)

var textLogger = logger.New()
var anoterTextLogger = logger.New("text")
var jsonLogger = logger.New("json")
```


## Creating a logger that streams to a remote FluentD 

> For the following to work, you'll need a FluentD service reachable
> at `localhost:24224`

```go
import (
  "github.com/usvc/go-log/pkg/logger"
  "github.com/usvc/go-log/pkg/hooks/fluentd"
  "github.com/usvc/go-log/pkg/constants"
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


## Releasing to GitHub

The GitHub URL for this repository is [https://github.com/usvc/go-log](https://github.com/usvc/go-log). The pipeline is configured to automatically push to this repository. Should the keys need to be regenerated, the `.ssh` Makefile recipe contains the commands required to generate the keys in a `.ssh` directory:

```sh
make .ssh
```

Inside the `.ssh` directory, copy the contents of `id_rsa.b64` and paste it as the `DEPLOY_KEY` CI/CD variable. Then copy the contents of `id_rsa.pub` and paste that as a deploy key with write access in the GitHub repository.


## Continuous integration/delivery (CI/CD) pipeline configuration

The following environment variables should be set in the CI/CD settings under Variables:

| Key | Description | Example |
| --- | --- | --- |
| `DEPLOY_KEY` | The base64 encoded private key that corresponds to the repository URL specified in `NEXT_REPO` | *(Output of `cat ~/.ssh/id_rsa \| base64 -w 0`)* |
| `NEXT_REPO_HOSTNAME` | The hostname of the `NEXT_REPO_URL` so that the domain's key can be verified | `github.com` |
| `NEXT_REPO_URL` | The SSH clone URL of the repository to push to in the `release` stage of the pipeline | `git@github.com:usvc/go-log.git` |

- - -

# License

This project is licensed under the [MIT license](./LICENSE).
