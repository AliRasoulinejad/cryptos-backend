package app

import (
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
	"github.com/AliRasoulinejad/cryptos-backend/internal/log"
)

func (application *Application) WithDB() {
	cfg := config.C.Database
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Logger.WithError(err).Fatal("error in connect to database")
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Logger.WithError(err).Fatal("error in connect open-telemetry to database")
	}

	application.DB = db
}
