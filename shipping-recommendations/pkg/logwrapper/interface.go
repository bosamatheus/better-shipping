package logwrapper

// Logger is the interface for the log wrapper.
type Logger interface {
	DefaultError(err error)
	InvalidArgError(argumentName string)
	InvalidArgValueError(argumentName, argumentValue string)
	MissingArgError(argumentName string)
}
