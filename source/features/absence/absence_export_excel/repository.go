package absenceexportexcel

import (
	"ayam_bangkok/source/features/absence/absence_export_excel/body"
	"context"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	getAbsenceForExport(ctx context.Context, start, end time.Time) ([]body.AbsenceExportRow, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
