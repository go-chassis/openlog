package openlogging_test

import (
	"github.com/go-chassis/paas-lager"
	"github.com/go-mesh/openlogging"
	"testing"
)

func TestGetLogger(t *testing.T) {
	log.Init(log.Config{
		LoggerLevel:   "DEBUG",
		LoggerFile:    "test.log",
		EnableRsyslog: false,
		LogFormatText: true,
		Writers:       []string{"file", "stdout"},
	})
	l := log.NewLogger("test")
	l.Infof("Hi %s, system is starting up ...", "paas-bot")
	openlogging.SetLogger(l)
	openlogging.GetLogger().Warn("test")
	openlogging.Info("test")
}
