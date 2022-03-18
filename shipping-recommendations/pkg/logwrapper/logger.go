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

// DefaultError is a standard generic error message.
func (l *StandardLogger) DefaultError(err error) {
	l.Errorf(genericErrorEvent.message, err)
}

// InvalidArgError is a standard invalid argument message.
func (l *StandardLogger) InvalidArgError(argumentName string) {
	l.Errorf(invalidArgEvent.message, argumentName)
}

// InvalidArgValueError is a standard invalid argument value message.
func (l *StandardLogger) InvalidArgValueError(argumentName, argumentValue string) {
	l.Errorf(invalidArgValueEvent.message, argumentName, argumentValue)
}

// MissingArgError is a standard missing argument message.
func (l *StandardLogger) MissingArgError(argumentName string) {
	l.Errorf(missingArgEvent.message, argumentName)
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
