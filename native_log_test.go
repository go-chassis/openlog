package openlog_test

import (
	"github.com/go-mesh/openlog"
	"testing"
)

func TestInfo(t *testing.T) {
	openlog.Info("test")
}
