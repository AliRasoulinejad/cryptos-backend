package log

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
)

var Logger *logrus.Logger

func InitLogger() {
	level, err := logrus.ParseLevel(config.C.Logger.Level)
	if err != nil {
		logrus.Fatalf("Error in initiating logger")
	}

	Logger = &logrus.Logger{
		Out:          os.Stderr,
		Level:        level,
		ReportCaller: true,
		Formatter: &logrus.TextFormatter{
			ForceColors:     true,
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			// LogFormat:       "[%time%] %lvl%, %msg%",
		},
	}
}
