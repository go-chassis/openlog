package openlogging_test

import (
	"testing"
	"github.com/go-mesh/openlogging"
)

func TestInfo(t *testing.T) {
	openlogging.Info("test")
}
