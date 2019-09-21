package fluentd

import (
	"strconv"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type UtilsTests struct {
	suite.Suite
}

func TestUtils(t *testing.T) {
	suite.Run(t, &UtilsTests{})
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
