package logs

import "github.com/sirupsen/logrus"

const (
	defaultTimestampFormat = "2006-01-02 15:04:05"
)

// Option is log option
type Option func(*logrus.Logger)

// NewLogger returns a logrus.Logger.
func NewLogger(options ...Option) *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: defaultTimestampFormat,
	})
	for _, optFunc := range options {
		optFunc(l)
	}

	return l
}

// WithLevel configures a logrus logger to log at level for all events.
func WithLevel(level logrus.Level) Option {
	return func(c *logrus.Logger) {
		c.Level = level
	}
}

// WithFormatter configures a logrus logger formatter
func WithFormatter(formatter logrus.Formatter) Option {
	return func(l *logrus.Logger) {
		l.SetFormatter(formatter)
	}
}
