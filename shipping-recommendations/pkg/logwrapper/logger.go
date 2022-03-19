package logwrapper

import (
	"time"

	"github.com/sirupsen/logrus"
)

// StandardLogger enforces specific log message formats.
type StandardLogger struct {
	*logrus.Logger
}

const (
	developmentEnv = "development"
	stagingEnv     = "staging"
	productionEnv  = "production"
)

// NewStandardLogger initializes a logrus-based standard logger according to the environment.
func NewStandardLogger(env string) *StandardLogger {
	baseLogger := logrus.New()
	standardLogger := StandardLogger{baseLogger}
	standardLogger.setLevel(env)
	standardLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	return &standardLogger
}

// InfoMsg is a standard info message.
func (l *StandardLogger) InfoMsg(msg string) {
	l.Infof(infoMsgEvent.msg, msg)
}

// UnexpectedErr is a standard unexpected error message.
func (l *StandardLogger) UnexpectedErr(err error) {
	l.Errorf(unexpectedErrEvent.msg, err)
}

// InvalidArgErr is a standard invalid argument message.
func (l *StandardLogger) InvalidArgErr(argName string) {
	l.Errorf(invalidArgErrEvent.msg, argName)
}

// InvalidArgValErr is a standard invalid argument value message.
func (l *StandardLogger) InvalidArgValErr(argName, argValue string) {
	l.Errorf(invalidArgValErrEvent.msg, argName, argValue)
}

// MissingArgErr is a standard missing argument message.
func (l *StandardLogger) MissingArgErr(argumentName string) {
	l.Errorf(missingArgErrEvent.msg, argumentName)
}

// setLevel sets the log level based on the environment.
func (l *StandardLogger) setLevel(env string) {
	switch env {
	case developmentEnv:
		l.SetLevel(logrus.DebugLevel)
	case stagingEnv:
		l.SetLevel(logrus.InfoLevel)
	case productionEnv:
		l.SetLevel(logrus.ErrorLevel)
	}
}
