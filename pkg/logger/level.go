package logger

// LogLevel is the type of log level
type LogLevel int8

const (
	// LogLevelTrace is the trace log level
	LogLevelTrace LogLevel = iota - 2
	// LogLevelDebug is the debug log level
	LogLevelDebug
	// LogLevelInfo is the info log level
	LogLevelInfo
	// LogLevelWarn is the warn log level
	LogLevelWarn
	// LogLevelError is the error log level
	LogLevelError
	// LogLevelFatal is the fatal log level
	LogLevelFatal
)

// String returns the string representation of the log level
func (l LogLevel) String() string {
	switch l {
	case LogLevelTrace:
		return "TRACE"
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}
