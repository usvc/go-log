package main

import (
	"github.com/usvc/go-log/pkg/logger"
)

var log = logger.New("json")

func main() {
	log.Infof("hello world")
}
