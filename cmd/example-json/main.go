package main

import (
	"gitlab.com/usvc/modules/go/log/pkg/logger"
)

var log = logger.New("json")

func main() {
	log.Infof("hello world")
}
