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
		&models.OrderMenuModel{},
	); err != nil {
		logger.Log.Fatal().Msg("Migration failed")
	}

	dbConn.Exec(`
		ALTER TABLE absences
		MODIFY clock_in TIME,
		MODIFY clock_out TIME NULL;
	`)

	logger.Log.Info().Msg("Migration success")

	// seedAbsence(dbConn)
	// logger.Log.Info().Msg("Seeding success")
}

// func seedAbsence(dbConn *gorm.DB) {
// 	logger.Log.Info().Msg("Seeding absence data...")

// 	userID := uint64(1)

// 	startDate := time.Now().AddDate(-1, 0, 0) // 1 tahun ke belakang
// 	endDate := time.Now()

// 	var absences []models.AbsenceModel

// 	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
// 		clockOut := "17:00:00"

// 		absences = append(absences, models.AbsenceModel{
// 			UserID:       userID,
// 			Date:         d,
// 			ClockIn:      "08:00:00",
// 			ClockOut:     &clockOut,
// 			LatitudeIn:   -6.200000,
// 			LongitudeIn:  106.816666,
// 			LatitudeOut:  -6.200000,
// 			LongitudeOut: 106.816666,
// 		})
// 	}

// 	if len(absences) > 0 {
// 		if err := dbConn.CreateInBatches(absences, 100).Error; err != nil {
// 			logger.Log.Fatal().Err(err).Msg("Failed to seed absences")
// 		}
// 	}

// 	logger.Log.Info().
// 		Int("total", len(absences)).
// 		Msg("Absence seeding completed")
// }
