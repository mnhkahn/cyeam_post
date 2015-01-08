package logs

import (
	"testing"
)

func TestMain(t *testing.T) {
	Log.Info("Test\n")
	Log.Trace("Hello\n")
	Log.Debug("World\n")
	Log.Close()
}
