package openlog_test

import (
	"github.com/go-chassis/openlog"
	"testing"
)

func TestInfo(t *testing.T) {
	openlog.Info("test")
}
