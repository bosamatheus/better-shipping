package logwrapper

// LogEvent represents a standard log event.
type LogEvent struct {
	msg string
}

var (
	unexpectedErrEvent    = LogEvent{"an unexpected error has occurred: %s"}
	infoMsgEvent          = LogEvent{"[INFO] %s"}
	invalidArgErrEvent    = LogEvent{"invalid arg: %s"}
	invalidArgValErrEvent = LogEvent{"invalid value for argument: %s: %v"}
	missingArgErrEvent    = LogEvent{"missing arg: %s"}
)
