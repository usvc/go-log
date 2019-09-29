package fluentd

import (
	"strconv"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/usvc/go-log/pkg/constants"
)

type UtilsTests struct {
	suite.Suite
}

func TestUtils(t *testing.T) {
	suite.Run(t, &UtilsTests{})
}

func (s *UtilsTests) Test_formatEntry() {
	inputEntry := logrus.Entry{
		Data: logrus.Fields{
			"field_one": "_field_one",
		},
		Level:   logrus.TraceLevel,
		Message: "__message",
	}
	formattedEntry := formatEntry(&inputEntry)
	s.Equal(map[string]interface{}{
		"field_one": "_field_one",
	}, formattedEntry[constants.FieldData])
	s.Equal("trace", formattedEntry[constants.FieldLevel])
	s.Equal("__message", formattedEntry[constants.FieldMessage])
}

func (s *UtilsTests) Test_hasReachedRetryLimit() {
	hookConfig := HookConfig{InitializeRetryCount: 1}
	hookExceeded := Hook{config: &hookConfig, retryCount: 2}
	hookNotExceeded := Hook{config: &hookConfig, retryCount: 1}
	s.Equal(true, hasReachedRetryLimit(&hookExceeded))
	s.Equal(false, hasReachedRetryLimit(&hookNotExceeded))
}

func (s *UtilsTests) Test_hasNoRetryLimit() {
	hookConfig := HookConfig{InitializeRetryCount: -1}
	hook := Hook{config: &hookConfig}
	s.Equal(true, hasNoRetryLimit(&hook))
}

func (s *UtilsTests) Test_spliceLogEntry() {
	expected := []*logrus.Entry{}
	for i := 0; i < 10; i++ {
		expected = append(expected, &logrus.Entry{Message: strconv.Itoa(i)})
	}
	s.Equal("2", expected[2].Message)
	expected = spliceLogEntry(expected, 2)
	s.Equal("3", expected[2].Message)
	s.Len(expected, 9)
}
