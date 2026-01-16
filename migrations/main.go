package main

import (
	"ayam_bangkok/source/common/models"
	"ayam_bangkok/source/config"
	"ayam_bangkok/source/pkg/db"
	"ayam_bangkok/source/pkg/logger"
)

func main() {
	logger.Init(true)
	
	cfg := config.Load()

	dbConn, err := db.Database(cfg)
	if err != nil {
		logger.Log.Fatal().Msg("Database connection failed")
	}
	logger.Log.Info().Msg("Running database migration...")

	if err := dbConn.AutoMigrate(
		&models.UserModel{},
		&models.RefreshTokenModel{},
		&models.AbsenceModel{},
		&models.MenuModel{},
		&models.RoomModel{},
		&models.FeedbackModel{},
	); err != nil {
		logger.Log.Fatal().Msg("Migration failed")
	}
	logger.Log.Info().Msg("Migration success")
}
