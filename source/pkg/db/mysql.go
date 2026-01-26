package db

import (
	"ayam_bangkok/source/config"
	"ayam_bangkok/source/pkg/logger"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error().Msg("Failed connect db")
		return nil, err
	}

	return db, nil
}