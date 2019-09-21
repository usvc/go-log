package fluentd

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type HookTests struct {
	suite.Suite
}

func TestHook(t *testing.T) {
	suite.Run(t, &HookTests{})
}

func (s *HookTests) TestLevels() {
	expectedLevels := []logrus.Level{
		logrus.TraceLevel,
	}
	hook := Hook{
		config: &HookConfig{
			Levels: expectedLevels,
		},
	}
	s.Equal(expectedLevels, hook.Levels())
}
