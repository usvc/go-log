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

func (s *HookTests) Test_getQueuedEntryAt() {
	hook := Hook{
		config: &HookConfig{},
		queue: []*logrus.Entry{
			&logrus.Entry{Message: "__queue_item_0"},
			&logrus.Entry{Message: "__queue_item_1"},
			&logrus.Entry{Message: "__queue_item_2"},
		},
	}
	s.Equal("__queue_item_0", hook.getQueuedEntryAt(0).Message)
	s.Equal("__queue_item_1", hook.getQueuedEntryAt(1).Message)
	s.Equal("__queue_item_2", hook.getQueuedEntryAt(2).Message)
}

func (s *HookTests) Test_removeLogFromQueue() {
	hook := Hook{
		config: &HookConfig{},
		queue: []*logrus.Entry{
			&logrus.Entry{Message: "__queue_item_0"},
			&logrus.Entry{Message: "__queue_item_1"},
			&logrus.Entry{Message: "__queue_item_2"},
		},
	}
	hook.removeLogFromQueue(1)
	s.Equal("__queue_item_0", hook.getQueuedEntryAt(0).Message)
	s.Equal("__queue_item_2", hook.getQueuedEntryAt(1).Message)
}

func (s *HookTests) Test_shouldRetry() {
	hookConfigLimited := HookConfig{InitializeRetryCount: 2}
	hookConfigUnlimited := HookConfig{InitializeRetryCount: -1}
	hookExceeded := Hook{config: &hookConfigLimited, retryCount: 3}
	hookNotExceeded := Hook{config: &hookConfigLimited, retryCount: 1}
	hookUnlimited := Hook{config: &hookConfigUnlimited}
	s.Equal(false, hookExceeded.shouldRetry())
	s.Equal(true, hookNotExceeded.shouldRetry())
	s.Equal(true, hookUnlimited.shouldRetry())
}
