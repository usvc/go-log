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
