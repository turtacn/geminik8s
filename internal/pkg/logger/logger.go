package logger

import (
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Logger defines a standard interface for logging.
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
}

// logrusLogger is a wrapper around logrus.Logger to implement the Logger interface.
type logrusLogger struct {
	*logrus.Entry
}

// NewLogger creates and configures a new logger.
func NewLogger(level string, output io.Writer, format string) Logger {
	log := logrus.New()

	log.SetOutput(output)

	// Set log level
	logLevel, err := logrus.ParseLevel(strings.ToLower(level))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	log.SetLevel(logLevel)

	// Set log format
	if strings.ToLower(format) == "json" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	return &logrusLogger{logrus.NewEntry(log)}
}

// NewDefaultLogger creates a logger with default settings.
func NewDefaultLogger() Logger {
	return NewLogger("info", os.Stdout, "text")
}

// WithField adds a single field to the log entry.
func (l *logrusLogger) WithField(key string, value interface{}) Logger {
	return &logrusLogger{l.Entry.WithField(key, value)}
}

// WithFields adds multiple fields to the log entry.
func (l *logrusLogger) WithFields(fields map[string]interface{}) Logger {
	return &logrusLogger{l.Entry.WithFields(fields)}
}

//Personal.AI order the ending
