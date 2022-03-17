package logwrapper

type Logger interface {
	DefaultError(err error)
	InvalidArgError(argumentName string)
	InvalidArgValueError(argumentName, argumentValue string)
	MissingArgError(argumentName string)
}
