package logger

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UtilsTest struct {
	suite.Suite
}

func TestUtilsSuite(t *testing.T) {
	suite.Run(t, &UtilsTest{})
}

func (s *UtilsTest) Test_callerPrettyfierSimplified() {
	observedFunc, observedFile := callerPrettyfierSimplified(&runtime.Frame{
		Line:     1337,
		Function: "path.to.function",
		File:     "path/to/file",
	})
	s.Equal("file:1337/function", observedFunc)
	s.Equal("", observedFile)
}

func (s *UtilsTest) Test_callerPrettyfier() {
	observedFunc, observedFile := callerPrettyfier(&runtime.Frame{
		Line:     1337,
		Function: "path.to.function",
		File:     "path/to/file",
	})
	s.Equal("file:1337", observedFile)
	s.Equal("function", observedFunc)
}

func (s *UtilsTest) Test_getFileAndLine() {
	s.Equal("file:1337", getFileAndLine(&runtime.Frame{
		File: "path/to/file",
		Line: 1337,
	}))
}

func (s *UtilsTest) Test_getFunctionBase() {
	s.Equal("function", getFunctionBase(&runtime.Frame{
		Function: "path.(toSomeOther).function",
	}))
}
