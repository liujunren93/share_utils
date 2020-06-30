package log

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	if Logger == nil {
		Logger = logrus.New()
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02T15:04:05.000",
		})
	}
}
