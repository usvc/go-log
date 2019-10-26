package fluentd

import (
	"runtime"
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

func (s *UtilsTests) Test_addCustomDataToLog() {
	entry := logrus.Entry{
		Data: logrus.Fields{
			"field1": "value1",
			"field2": 2,
		},
	}
	log := map[string]interface{}{}
	addCustomDataToLog(log, &entry)
	s.Equal(map[string]interface{}{
		"field1": "value1",
		"field2": 2,
	}, log[constants.FieldData])
}

func (s *UtilsTests) Test_addCallerDataToLog() {
	entry := logrus.Entry{
		Caller: &runtime.Frame{
			File:     "__caller_file",
			Function: "__caller_function",
			Line:     12345,
		},
	}
	log := map[string]interface{}{}
	addCallerDataToLog(log, &entry)
	s.Equal("__caller_file:12345", log[constants.FieldFile])
	s.Equal("__caller_function", log[constants.FieldFunction])
}

func (s *UtilsTests) Test_clearQueue() {
	mock := hookMock{
		instance: &Hook{
			queue: []*logrus.Entry{
				&logrus.Entry{
					Message: "__message_1",
				},
				&logrus.Entry{
					Message: "__message_2",
				},
			},
		},
	}
	clearQueue(&mock)
	s.Equal(uint(2), mock._SendCount)
	s.Equal("__message_1", mock._SendArgs[0][constants.FieldMessage])
	s.Equal("__message_2", mock._SendArgs[1][constants.FieldMessage])
}

func (s *UtilsTests) Test_createLogFromEntry_noCaller() {
	inputEntry := logrus.Entry{
		Data:    logrus.Fields{"field_one": "_field_one"},
		Level:   logrus.TraceLevel,
		Message: "__message",
	}
	formattedEntry := createLogFromEntry(&inputEntry)

	s.NotNil(formattedEntry[constants.FieldData],
		"'%s' (constants.FieldData) field exists", constants.FieldData)
	s.Equal(map[string]interface{}{
		"field_one": "_field_one",
	}, formattedEntry[constants.FieldData])

	s.NotNil(formattedEntry[constants.FieldLevel],
		"'%s' (constants.FieldLevel) field exists", constants.FieldLevel)
	s.Equal("trace", formattedEntry[constants.FieldLevel])

	s.NotNil(formattedEntry[constants.FieldMessage],
		"'%s' (constants.FieldMessage) field exists", constants.FieldMessage)
	s.Equal("__message", formattedEntry[constants.FieldMessage])

	s.Nil(formattedEntry[constants.FieldFile],
		"'%s' (constants.FieldFile) should not exist when not passed in", constants.FieldFile)
	s.Nil(formattedEntry[constants.FieldFunction],
		"'%s' (constants.FieldFunction) should not exist when not passed in", constants.FieldFunction)
}

func (s *UtilsTests) Test_createLogFromEntry_withCaller() {
	inputEntry := logrus.Entry{
		Data:    logrus.Fields{"field_one": "_field_one"},
		Level:   logrus.TraceLevel,
		Message: "__message",
		Caller: &runtime.Frame{
			File:     "__caller_file",
			Line:     12345,
			Function: "__caller_function",
		},
	}
	formattedEntry := createLogFromEntry(&inputEntry)

	s.NotNil(formattedEntry[constants.FieldFile],
		"'%s' (constants.FieldFile) field exists", constants.FieldFile)
	s.Equal("__caller_file:12345", formattedEntry[constants.FieldFile])

	s.NotNil(formattedEntry[constants.FieldFunction],
		"'%s' (constants.FieldFunction) field exists", constants.FieldFunction)
	s.Equal("__caller_function", formattedEntry[constants.FieldFunction])
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
