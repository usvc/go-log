package logger

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type ConfigTest struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, &ConfigTest{})
}

func (c *ConfigTest) Test_configureLogger() {
	logger := &logrus.Logger{}
	configureLogger(logger)
	c.Equal(logger.Out, os.Stdout)
	c.Equal(logger.Level, logrus.TraceLevel)
	c.True(logger.ReportCaller)
}
