package openlog_test

import (
	"errors"
	"github.com/go-chassis/openlog"
	"testing"
)

func TestGetLogger(t *testing.T) {
	openlog.Warn("depth")
	openlog.Warn("test", openlog.WithTags(openlog.Tags{
		"hi":  "asd",
		"asd": "asd",
	}), openlog.WithErr(errors.New("er")))
	openlog.Info("test")
}
