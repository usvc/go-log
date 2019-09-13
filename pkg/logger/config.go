package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"gitlab.com/usvc/modules/go/logging/pkg/constants"
)

var TextFormatter = &logrus.TextFormatter{
	CallerPrettyfier: func(r *runtime.Frame) (string, string) {
		return fmt.Sprintf("%s/%s", path.Base(r.File), path.Base(strings.Replace(r.Function, ".", "/", -1))), ""
	},
	DisableSorting:   false,
	FieldMap:         constants.DefaultFieldMap,
	FullTimestamp:    true,
	TimestampFormat:  constants.TimestampFormat,
	QuoteEmptyFields: true,
}

var JSONFormatter = &logrus.JSONFormatter{
	CallerPrettyfier: func(r *runtime.Frame) (string, string) {
		return fmt.Sprintf("%s/%s", path.Base(r.File), path.Base(strings.Replace(r.Function, ".", "/", -1))), ""
	},
	DataKey:         constants.FieldData,
	FieldMap:        constants.DefaultFieldMap,
	TimestampFormat: constants.TimestampFormat,
}
