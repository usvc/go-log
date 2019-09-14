package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func callerPrettyfierSimplified(r *runtime.Frame) (string, string) {
	return fmt.Sprintf("%s/%s", getFileAndLine(r), getFunctionBase(r)), ""
}

func callerPrettyfier(r *runtime.Frame) (string, string) {
	return getFunctionBase(r), getFileAndLine(r)
}

func getFileAndLine(runtimeFrame *runtime.Frame) string {
	return fmt.Sprintf("%s:%v", path.Base(runtimeFrame.File), runtimeFrame.Line)
}

func getFunctionBase(runtimeFrame *runtime.Frame) string {
	return path.Base(strings.Replace(runtimeFrame.Function, ".", "/", -1))
}
