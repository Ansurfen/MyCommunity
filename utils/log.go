package utils

import (
	"github.com/sirupsen/logrus"
)

func init() {
	formatter := &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		DisableQuote:    true,
		TimestampFormat: Conf.GetString("logger.format.timestamp"),
	}
	logrus.SetFormatter(formatter)
	logrus.Debug(Conf.GetString("logger.level"))
}
