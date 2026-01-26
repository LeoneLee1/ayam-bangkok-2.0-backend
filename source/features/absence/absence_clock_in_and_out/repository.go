package absenceclockinandout

import (
	"ayam_bangkok/source/common/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	getAbsenceToday(ctx context.Context, userID uint, today time.Time) (*models.AbsenceModel, error)
	absenceIn(ctx context.Context, userID uint, absence *models.AbsenceModel) error
	absenceOut(ctx context.Context, userID uint, absence *models.AbsenceModel, today time.Time) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
