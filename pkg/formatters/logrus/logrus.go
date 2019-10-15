package logrus

import (
	liblogrus "github.com/sirupsen/logrus"
	"github.com/usvc/go-log/pkg/constants"
)

var Text = &liblogrus.TextFormatter{
	CallerPrettyfier: callerPrettyfierSimplified,
	DisableSorting:   false,
	FieldMap:         constants.DefaultFieldMap,
	FullTimestamp:    true,
	QuoteEmptyFields: true,
	TimestampFormat:  constants.TimestampFormat,
}

var JSON = &liblogrus.JSONFormatter{
	CallerPrettyfier: callerPrettyfier,
	DataKey:          constants.FieldData,
	FieldMap:         constants.DefaultFieldMap,
	TimestampFormat:  constants.TimestampFormat,
}
