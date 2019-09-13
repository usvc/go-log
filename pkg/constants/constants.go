package constants

import (
	"github.com/sirupsen/logrus"
)

var DefaultFieldMap = logrus.FieldMap{
	logrus.FieldKeyTime:  FieldTimestamp,
	logrus.FieldKeyMsg:   FieldMessage,
	logrus.FieldKeyLevel: FieldLevel,
	logrus.FieldKeyFunc:  FieldFunction,
	logrus.FieldKeyFile:  FieldFile,
}

const FieldData = "@data"
const FieldFile = "@file"
const FieldFunction = "@function"
const FieldLevel = "@level"
const FieldMessage = "@message"
const FieldTimestamp = "@timestamp"

const TimestampFormat = "2006-01-02T15:04:05"
