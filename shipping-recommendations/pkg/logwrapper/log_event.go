package logwrapper

// LogEvent represents a standard log event
type LogEvent struct {
	message string
}

var (
	genericErrorEvent    = LogEvent{"an error has occurred: %s"}
	invalidArgEvent      = LogEvent{"invalid arg: %s"}
	invalidArgValueEvent = LogEvent{"invalid value for argument: %s: %v"}
	missingArgEvent      = LogEvent{"missing arg: %s"}
)
