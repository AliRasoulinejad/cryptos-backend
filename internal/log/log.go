package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"

	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
)

var Logger *logrus.Logger

func InitLogger() {
	level, err := logrus.ParseLevel(config.C.Logger.Level)
	if err != nil {
		logrus.Fatalf("Error in initiating logger")
	}

	Logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: level,
		Hooks: make(logrus.LevelHooks),
		// ReportCaller: true,
		Formatter: &logrus.TextFormatter{
			ForceColors:            true,
			DisableColors:          false,
			DisableLevelTruncation: true,
			TimestampFormat:        "2006-01-02 15:04:05",
			// LogFormat:       "[%time%] %lvl%, %msg%",
		},
	}

	Logger.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
	)))
}
