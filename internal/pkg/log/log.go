package log

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log = logrus.New()

func init() {
	log.Formatter = defaultFormatter()
	log.Level = logrus.DebugLevel
}

func defaultFormatter() *prefixed.TextFormatter {
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true

	// Set specific colors for prefix and timestamp
	formatter.SetColorScheme(&prefixed.ColorScheme{
		PrefixStyle:    "blue+b",
		TimestampStyle: "white+h",
	})

	return formatter
}

func Print(instance, msg string) {
	if instance != "" {
		log.WithFields(logrus.Fields{
			"prefix": instance,
		}).Info(msg)
	} else {
		log.Info(msg)
	}
}

func Fatal(instance, msg string, err error) {
	log.WithFields(logrus.Fields{
		"prefix": instance,
	}).Fatal(msg, ": ", err.Error())
}

func Error(instance, msg string, err error) {
	log.WithFields(logrus.Fields{
		"prefix": instance,
	}).Errorln(msg, ": ", err.Error())
}
