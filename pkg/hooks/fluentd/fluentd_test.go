package fluentd

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type FluentDTests struct {
	suite.Suite
}

func TestFluentD(t *testing.T) {
	suite.Run(t, &FluentDTests{})
}

func (s *FluentDTests) TestNewHook_createsQueue() {
	hook := NewHook(&HookConfig{})
	s.NotNil(hook.queue)
}

func (s *FluentDTests) TestNewHook_usesDefaultLog() {
	hook := NewHook(&HookConfig{})
	expectedLogPointer := fmt.Sprintf("%p", log)
	s.Equal(expectedLogPointer, fmt.Sprintf("%p", hook.log))
}

func (s *FluentDTests) TestNewHook_usesInputConfig() {
	config := HookConfig{}
	expectedConfigPointer := fmt.Sprintf("%p", &config)
	hook := NewHook(&config)
	s.Equal(expectedConfigPointer, fmt.Sprintf("%p", hook.config))
}

func (s *FluentDTests) TestNewHook_usesInputLog() {
	logger := &logrus.Logger{}
	hook := NewHook(&HookConfig{}, logger)
	defaultLogPointer := fmt.Sprintf("%p", log)
	expectedLogPointer := fmt.Sprintf("%p", logger)
	s.NotEqual(defaultLogPointer, fmt.Sprintf("%p", hook.log))
	s.Equal(expectedLogPointer, fmt.Sprintf("%p", hook.log))
}
