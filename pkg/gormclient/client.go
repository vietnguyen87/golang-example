package gormclient

import (
	"context"
	"example-service/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"example-service/pkg/logger"
)

func New() (*gorm.DB, error) {
	log := logger.CToL(context.Background(), "gormclient.New")

	cfg := config.GetDatabaseConfig()

	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Errorf("gorm.Open returns error: %s", err.Error())
		return nil, err
	}

	return db, nil
}
