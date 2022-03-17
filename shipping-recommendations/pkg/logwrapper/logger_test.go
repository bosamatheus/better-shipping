package logwrapper

import (
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestNewStandardLogger(t *testing.T) {
	testCases := []struct {
		name string
		env  string
		want *StandardLogger
	}{
		{
			"DevelopmentEnvironment",
			developmentEnv,
			&StandardLogger{&logrus.Logger{Level: logrus.DebugLevel}},
		},
		{
			"StagingEnvironment",
			stagingEnv,
			&StandardLogger{&logrus.Logger{Level: logrus.InfoLevel}},
		},
		{
			"ProductionEnvironment",
			productionEnv,
			&StandardLogger{&logrus.Logger{Level: logrus.ErrorLevel}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewStandardLogger(tc.env)

			assert.Equal(t, tc.want.GetLevel(), got.GetLevel())
		})
	}
}

func TestStandardLogger_DefaultError(t *testing.T) {
	nullLogger, hook := test.NewNullLogger()
	standardLogger := StandardLogger{nullLogger}

	standardLogger.DefaultError(errors.New("generic error"))

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "an error has occurred: generic error", hook.LastEntry().Message)
}

func TestStandardLogger_InvalidArgError(t *testing.T) {
	nullLogger, hook := test.NewNullLogger()
	standardLogger := StandardLogger{nullLogger}

	standardLogger.InvalidArgError("argumentName")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "invalid arg: argumentName", hook.LastEntry().Message)
}

func TestStandardLogger_InvalidArgValueError(t *testing.T) {
	nullLogger, hook := test.NewNullLogger()
	standardLogger := StandardLogger{nullLogger}

	standardLogger.InvalidArgValueError("argumentName", "argumentValue")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "invalid value for argument: argumentName: argumentValue", hook.LastEntry().Message)
}

func TestStandardLogger_MissingArgError(t *testing.T) {
	nullLogger, hook := test.NewNullLogger()
	standardLogger := StandardLogger{nullLogger}

	standardLogger.MissingArgError("argumentName")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "missing arg: argumentName", hook.LastEntry().Message)
}
