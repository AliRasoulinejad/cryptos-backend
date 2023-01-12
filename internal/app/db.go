package app

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
)

func (application *Application) WithDB() {
	cfg := config.C.Database
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.WithError(err).Fatal("error in connect to database")
	}

	application.DB = db
}
