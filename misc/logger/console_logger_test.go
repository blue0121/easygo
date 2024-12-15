package logger

import "testing"

func TestOutput(t *testing.T) {
	SetLogger(&consoleLogger{})
	Debug("Write file")
	Info("Write file")
	Warn("Write file")
	Error("Write file")
}
