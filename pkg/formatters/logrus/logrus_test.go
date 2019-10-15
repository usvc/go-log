package logrus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"testing"

	liblogrus "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type LogrusTests struct {
	expectedTimestampRegexp string
	expectedFieldData       liblogrus.Fields
	output                  bytes.Buffer
	logger                  *liblogrus.Logger
	suite.Suite
}

func TestLogruses(t *testing.T) {
	suite.Run(t, &LogrusTests{})
}

func (s *LogrusTests) SetupTest() {
	s.expectedTimestampRegexp = "[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}"
	s.expectedFieldData = liblogrus.Fields{
		"__test_string": "__test",
		"__test_char":   '_',
		"__test_int":    1234567,
		"__test_uint":   uint(1234567),
		"__test_float":  1234.567,
		"__test_bool":   false,
	}
	s.logger = liblogrus.New()
	s.logger.SetLevel(liblogrus.TraceLevel)
	s.logger.SetOutput(&s.output)
	s.logger.SetReportCaller(true)
}

func (s *LogrusTests) TestTimestampFormatSame() {
	s.logger.SetFormatter(JSON)
	s.logger.Trace("json")
	s.logger.SetFormatter(Text)
	s.logger.Trace("text")
	log := s.output.String()
	s.output.Reset()
	re := regexp.MustCompile(fmt.Sprintf("(%s)", s.expectedTimestampRegexp))
	match := re.FindStringSubmatch(log)
	s.Len(match, 2)
	s.Equal(match[0], match[1])
}

func (s *LogrusTests) TestText() {
	expectedFunctionSignature := `[a-zA-Z0-9_-]+.go:[0-9]+/TestText`
	expectedMessage := "__test_text"
	s.logger.SetFormatter(Text)
	s.logger.WithFields(s.expectedFieldData).Trace(expectedMessage)
	log := s.output.String()
	s.output.Reset()
	s.Regexp(fmt.Sprintf(`@timestamp="%s"`, s.expectedTimestampRegexp), log)
	s.Contains(log, fmt.Sprintf("@message=%s", expectedMessage))
	s.Contains(log, "@level=trace")
	s.Regexp(fmt.Sprintf(`@function="%s"`, expectedFunctionSignature), log)
	s.Contains(log, fmt.Sprintf("__test_string=%s", s.expectedFieldData["__test_string"]))
	s.Contains(log, fmt.Sprintf("__test_char=%v", s.expectedFieldData["__test_char"]))
	s.Contains(log, fmt.Sprintf("__test_bool=%v", s.expectedFieldData["__test_bool"]))
	s.Contains(log, fmt.Sprintf("__test_int=%v", s.expectedFieldData["__test_int"]))
	s.Contains(log, fmt.Sprintf("__test_uint=%v", s.expectedFieldData["__test_uint"]))
	s.Contains(log, fmt.Sprintf("__test_float=%v", s.expectedFieldData["__test_float"]))
}

func (s *LogrusTests) TestJSON() {
	expectedFileSignature := `[a-zA-Z0-9_-]+.go:[0-9]+`
	expectedMessage := "__test_json"
	s.logger.SetFormatter(JSON)
	s.logger.WithFields(s.expectedFieldData).Trace(expectedMessage)
	log := s.output.String()
	s.output.Reset()
	var structuredLog map[string]interface{}
	err := json.Unmarshal([]byte(log), &structuredLog)
	if !s.Nil(err) {
		panic(err)
	}
	s.Regexp(s.expectedTimestampRegexp, structuredLog["@timestamp"])
	s.Equal(expectedMessage, structuredLog["@message"])
	s.Equal("trace", structuredLog["@level"])
	s.Regexp("TestJSON", structuredLog["@function"])
	s.Regexp(expectedFileSignature, structuredLog["@file"])
	if data, ok := structuredLog["@data"].(map[string]interface{}); ok {
		s.Equal(s.expectedFieldData["__test_string"], data["__test_string"])
		if charValue, ok := data["__test_char"].(int32); ok {
			s.Equal(s.expectedFieldData["__test_char"], charValue)
		}
		s.Equal(s.expectedFieldData["__test_bool"], data["__test_bool"])
		if intValue, ok := data["__test_int"].(int); ok {
			s.Equal(s.expectedFieldData["__test_int"], intValue)
		}
		if uintValue, ok := data["__test_uint"].(uint); ok {
			s.Equal(s.expectedFieldData["__test_uint"], uintValue)
		}
		s.Equal(s.expectedFieldData["__test_float"], data["__test_float"])
	}
}
