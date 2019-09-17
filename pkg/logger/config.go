package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/usvc/go-log/pkg/constants"
)

var TextFormatter = &logrus.TextFormatter{
	CallerPrettyfier: callerPrettyfierSimplified,
	DisableSorting:   false,
	FieldMap:         constants.DefaultFieldMap,
	FullTimestamp:    true,
	QuoteEmptyFields: true,
	TimestampFormat:  constants.TimestampFormat,
}

var JSONFormatter = &logrus.JSONFormatter{
	CallerPrettyfier: callerPrettyfier,
	DataKey:          constants.FieldData,
	FieldMap:         constants.DefaultFieldMap,
	TimestampFormat:  constants.TimestampFormat,
}

func configureLogger(logger *logrus.Logger) {
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.TraceLevel)
	logger.SetReportCaller(true)
}
