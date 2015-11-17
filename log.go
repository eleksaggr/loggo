package loggo

import (
	"fmt"
  "image/color"
	"os"
)

type Logger struct {
	debugMode bool
	output    *os.File
}

func NewLogger(output *os.File, debugMode bool) *Logger {
	logger := new(Logger)

	if output != nil {
		logger.output = output
	} else {
		// Print a warning to the user
		fmt.Println("[LOGGO] ---Falling back to Stderr---")
		logger.output = os.Stderr
	}
	logger.debugMode = debugMode

	return logger
}

func (logger *Logger) Log(section, msg string) {
	fmt.Fprintf(logger.output, "[%s] %s\n", section, msg)
}

func (logger *Logger) Debug(section, msg string) {
	if logger.debugMode == true {
		logger.Log(section, msg)
	}
}

func (logger *Logger) SetDebug(flag bool) {
	logger.debugMode = flag
}
