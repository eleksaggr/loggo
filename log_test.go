package loggo

import (
	"os"
	"testing"
)

func TestInfo(t *testing.T) {
	l := NewLog(os.Stdout)
	l.Info("This is a test.")
}

func TestDebug(t *testing.T) {
	l := NewLog(os.Stdout)
	l.Debug("This is a debug message.")
}

func TestWarning(t *testing.T) {
	l := NewLog(os.Stdout)
	l.Warning("This is a Warning.")
}

func TestError(t *testing.T) {
	l := NewLog(os.Stdout)
	l.Error("This shouldnt have happened!")
}
