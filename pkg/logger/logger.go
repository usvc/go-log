package logger

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/usvc/modules/go/log/lib/utils"
)

func New(formats ...string) *logrus.Logger {
	format := utils.ParseVariadicString(formats, "text")
	logger := logrus.New()
	configureLogger(logger)
	switch format {
	case "json":
		logger.SetFormatter(JSONFormatter)
	case "text":
		fallthrough
	default:
		logger.SetFormatter(TextFormatter)
	}
	return logger
}
