package fluentd

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/usvc/go-log/pkg/constants"
)

type ConfigTest struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, &ConfigTest{})
}

func (s *ConfigTest) Test_createFluentConfig_defaults() {
	fluentConfig := createFluentConfig(&HookConfig{})
	s.Equal(constants.DefaultFluentDHost, fluentConfig.FluentHost)
	s.Equal(constants.DefaultFluentDPort, fluentConfig.FluentPort)
	s.Equal(constants.DefaultFluentDTag, fluentConfig.TagPrefix)
	s.Equal(constants.DefaultFluentDTimeout, fluentConfig.Timeout)
	s.Equal(constants.DefaultFluentDTimeout, fluentConfig.WriteTimeout)
}

func (s *ConfigTest) Test_createFluentConfig_nonDefaults() {
	expectedHost := "host"
	expectedPort := 12345
	expectedTag := "tag"
	fluentConfig := createFluentConfig(&HookConfig{
		Host: expectedHost,
		Port: expectedPort,
		Tag:  expectedTag,
	})
	s.Equal(expectedHost, fluentConfig.FluentHost)
	s.Equal(expectedPort, fluentConfig.FluentPort)
	s.Equal(expectedTag, fluentConfig.TagPrefix)
}
