package logger

// Logger is a simple logger interface
type Logger interface {
	// Trace logs a message at level Trace on the logger.
	Trace(args ...any)
	// Tracef logs a message at level Trace on the logger.
	Tracef(format string, args ...any)
	// Debug logs a message at level Debug on the logger.
	Debug(args ...any)
	// Debugf logs a message at level Debug on the logger.
	Debugf(format string, args ...any)
	// Info logs a message at level Info on the logger.
	Info(args ...any)
	// Infof logs a message at level Info on the logger.
	Infof(format string, args ...any)
	// Warn logs a message at level Warn on the logger.
	Warn(args ...any)
	// Warnf logs a message at level Warn on the logger.
	Warnf(format string, args ...any)
	// Error logs a message at level Error on the logger.
	Error(args ...any)
	// Errorf logs a message at level Error on the logger.
	Errorf(format string, args ...any)
	// Fatal logs a message at level Fatal on the logger.
	Fatal(args ...any)
	// Fatalf logs a message at level Fatal on the logger.
	Fatalf(format string, args ...any)
}
