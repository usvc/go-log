package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func callerPrettyfier(r *runtime.Frame) (string, string) {
	return fmt.Sprintf("%s/%s", path.Base(r.File), path.Base(strings.Replace(r.Function, ".", "/", -1))), ""
}
