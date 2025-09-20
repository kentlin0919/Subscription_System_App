package log

import (
	stdlog "log"
	"os"
)

// Logger defines the logging contract for the application.
type Logger interface {
	Printf(format string, v ...any)
}

// StandardLogger delegates logging to the stdlib logger.
type StandardLogger struct {
	logger *stdlog.Logger
}

// NewStandardLogger creates a StandardLogger that writes to stdout.
func NewStandardLogger() *StandardLogger {
	return &StandardLogger{logger: stdlog.New(os.Stdout, "", stdlog.LstdFlags)}
}

// Printf logs formatted messages.
func (l *StandardLogger) Printf(format string, v ...any) {
	if l == nil || l.logger == nil {
		return
	}

	l.logger.Printf(format, v...)
}
