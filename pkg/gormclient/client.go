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
	log.Infof("config: %v", cfg.Dsn)

	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Errorf("gorm.Open config: %s returns error: %s", cfg.Dsn, err.Error())
		return nil, err
	}

	return db, nil
}
