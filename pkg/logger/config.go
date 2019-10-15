package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func configureLogger(logger *logrus.Logger) {
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.TraceLevel)
	logger.SetReportCaller(true)
}
