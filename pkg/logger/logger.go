package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func New(format string) *logrus.Logger {
	switch format {
	case "json":
		return NewJSONFormattedLogger()
	case "text":
		fallthrough
	default:
		return NewTextFormattedLogger()
	}
}

func NewTextFormattedLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.TraceLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(TextFormatter)
	return logger
}

func NewJSONFormattedLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.TraceLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(JSONFormatter)
	return logger
}
