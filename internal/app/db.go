package app

import (
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

	application.DB = db
}
