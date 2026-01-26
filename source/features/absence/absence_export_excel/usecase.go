package absenceexportexcel

import (
	"ayam_bangkok/source/features/absence/absence_export_excel/body"
	"context"
	"time"
)

type Usecase interface {
	getAbsenceForExport(ctx context.Context, start, end time.Time) ([]body.AbsenceExportRow, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
