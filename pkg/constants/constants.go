package constants

import (
	"time"

	"github.com/sirupsen/logrus"
)

// DefaultFieldMap sets the FieldMap property of the TextFormatter and
// JSONFormatter classes that define how logrus outputs logs
var DefaultFieldMap = logrus.FieldMap{
	logrus.FieldKeyTime:  FieldTimestamp,
	logrus.FieldKeyMsg:   FieldMessage,
	logrus.FieldKeyLevel: FieldLevel,
	logrus.FieldKeyFunc:  FieldFunction,
	logrus.FieldKeyFile:  FieldFile,
}

const DefaultFluentDHost = "localhost"
const DefaultFluentDPort = 24224
const DefaultFluentDTag = "app"
const DefaultFluentDTimeout = 5 * time.Second

// DefaultHookLevels can be used to set the levels for which the logrus
// hook will be activated if you'd like the hook to be activated on all
// levels of logs
var DefaultHookLevels = []logrus.Level{
	logrus.TraceLevel,
	logrus.DebugLevel,
	logrus.InfoLevel,
	logrus.WarnLevel,
	logrus.ErrorLevel,
	logrus.PanicLevel,
	logrus.FatalLevel,
}

// FieldData defines what the data field should be tagged
// (applies only for JSONFormatter)
const FieldData = "@data"

// FieldFile defines the key for the value containing the
// file name and line
const FieldFile = "@file"

// FieldFunction defines the key for the value containing the
// function name
const FieldFunction = "@function"

// FieldLevel defines the key for the value containing the level
// of the log entry
const FieldLevel = "@level"

// FieldMessage defines the key for the value containing the
// message of the log entry
const FieldMessage = "@message"

// FieldTimestamp defines the key for the value containing the
// timestamp of the log entry
const FieldTimestamp = "@timestamp"

// TimestampFormat defines the format which all logs should be
// in within the usvc namespace
const TimestampFormat = "2006-01-02T15:04:05"
