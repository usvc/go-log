package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/usvc/go-log/lib/utils"
	formatters "github.com/usvc/go-log/pkg/formatters/logrus"
)

func New(formats ...string) *logrus.Logger {
	format := utils.ParseVariadicString(formats, "text")
	logger := logrus.New()
	configureLogger(logger)
	switch format {
	case "json":
		logger.SetFormatter(formatters.JSON)
	case "text":
		fallthrough
	default:
		logger.SetFormatter(formatters.Text)
	}
	return logger
}
