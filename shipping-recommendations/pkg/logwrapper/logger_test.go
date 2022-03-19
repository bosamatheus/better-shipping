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

func TestStandardLogger_InfoMsg(t *testing.T) {
	nullLogger, hook := test.NewNullLogger()
	standardLogger := StandardLogger{nullLogger}

	standardLogger.InfoMsg("msg")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "[INFO] msg", hook.LastEntry().Message)
}

func TestStandardLogger_UnexpectedErr(t *testing.T) {
	nullLogger, hook := test.NewNullLogger()
	standardLogger := StandardLogger{nullLogger}

	standardLogger.UnexpectedErr(errors.New("error"))

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "an unexpected error has occurred: error", hook.LastEntry().Message)
}

func TestStandardLogger_InvalidArgErr(t *testing.T) {
	nullLogger, hook := test.NewNullLogger()
	standardLogger := StandardLogger{nullLogger}

	standardLogger.InvalidArgErr("argName")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "invalid arg: argName", hook.LastEntry().Message)
}

func TestStandardLogger_InvalidArgValErr(t *testing.T) {
	nullLogger, hook := test.NewNullLogger()
	standardLogger := StandardLogger{nullLogger}

	standardLogger.InvalidArgValErr("argName", "argVal")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "invalid value for argument: argName: argVal", hook.LastEntry().Message)
}

func TestStandardLogger_MissingArgErr(t *testing.T) {
	nullLogger, hook := test.NewNullLogger()
	standardLogger := StandardLogger{nullLogger}

	standardLogger.MissingArgErr("argName")

	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "missing arg: argName", hook.LastEntry().Message)
}
