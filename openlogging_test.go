package openglog_test

import (
	"github.com/go-mesh/openglog"
	"testing"
)

func TestGetLogger(t *testing.T) {
	openglog.Warn("depth")
	openglog.GetLogger().Warn("test", openglog.WithTags(openglog.Tags{
		"hi":  "asd",
		"asd": "asd",
	}))
	openglog.Info("test")
}
