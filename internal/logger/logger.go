package logger

import (
	"github.com/sirupsen/logrus"
)

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

func NewLogger(verbose bool) *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}
	if verbose {
		standardLogger.SetLevel(logrus.DebugLevel)
	} else {
		standardLogger.SetLevel(logrus.InfoLevel)
	}

	return standardLogger
}

func Debug(msg string) {
	logrus.Debug(msg)
}

func Warn(msg string) {
	logrus.Warn(msg)
}

func Error(msg string) {
	logrus.Error(msg)
}
