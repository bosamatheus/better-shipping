package logwrapper

type Logger interface {
	DefaultError(err error)
	InvalidArgError(argumentName string)
	InvalidArgValueError(argumentName string, argumentValue string)
	MissingArgError(argumentName string)
}
