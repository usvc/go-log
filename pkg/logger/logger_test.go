package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoggerTest struct {
	suite.Suite
}

func TestLoggerSuite(t *testing.T) {
	suite.Run(t, &LoggerTest{})
}

func (c *LoggerTest) Test_New_Text() {
	var log bytes.Buffer
	textLogger := New()
	textLogger.SetOutput(&log)
	textLogger.Print("hello world")
	c.Contains(log.String(), `@level=info @message="hello world"`)
}

func (c *LoggerTest) Test_New_JSON() {
	var log bytes.Buffer
	textLogger := New("json")
	textLogger.SetOutput(&log)
	textLogger.Print("hello world")
	c.Contains(log.String(), `"@message":"hello world"`)
	c.Contains(log.String(), `"@level":"info"`)
}
