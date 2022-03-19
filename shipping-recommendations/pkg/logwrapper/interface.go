package logwrapper

// Logger is the interface for the log wrapper.
type Logger interface {
	InfoMsg(msg string)
	UnexpectedErr(err error)
	InvalidArgErr(argName string)
	InvalidArgValErr(argName, argVal string)
	MissingArgErr(argName string)
}
