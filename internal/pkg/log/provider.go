package log

import "github.com/sirupsen/logrus"

func NewLogger(instance string) *logrus.Entry {
	newLogger := logrus.New()
	formatter := defaultFormatter()
	newLogger.Level = logrus.DebugLevel

	entry := newLogger.WithFields(logrus.Fields{
		"prefix": instance,
	})

	entry.Logger.Formatter = formatter

	return entry
}
