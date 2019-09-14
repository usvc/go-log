/*
Package logger is a standardised logger for use in usvc projects. Feel free to use it
in your own projects!

Basic text logger:

	```go
	var log = logger.New()
	log.Infof("hello world")
	```

Example output:

	INFO[2019-09-14T14:24:34]main.go:10/main hello world

Basic JSON logger:

	```go
	var log = logger.New()
	log.Infof("hello world")
	```

Example output:

	{"@data":{},"@file":"main.go:10","@function":"main","@level":"info","@message":"hello world","@timestamp":"2019-09-14T14:24:31"}

For the full guide visit https://gitlab.com/usvc/modules/go/log
*/
package logger
