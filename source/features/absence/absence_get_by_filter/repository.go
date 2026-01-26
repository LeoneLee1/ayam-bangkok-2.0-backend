package absencegetbyfilter

import (
	"ayam_bangkok/source/features/absence/absence_get_by_filter/body"
	"context"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	getAbsenceByFilter(ctx context.Context, start, end time.Time, page, limit int) ([]body.AbsenceFilterResponse, int64, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
