package loggo

import (
	// "os"
	"testing"
)

func TestLog(t *testing.T) {
	l := NewLogger(nil, false)
	l.Log("Test", "This is a test.")
}
