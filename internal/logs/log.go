package logs

import (
	"github.com/sirupsen/logrus"
)

type Service struct{}

func NewLogService() *Service {
	return &Service{}
}

func InitLog(format, level string) {
	switch level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	if format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}
