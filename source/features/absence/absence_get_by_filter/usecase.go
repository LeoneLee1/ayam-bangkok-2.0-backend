package absencegetbyfilter

import (
	"ayam_bangkok/source/features/absence/absence_get_by_filter/body"
	"context"
	"time"
)

type Usecase interface {
	getAbsenceByFilter(ctx context.Context, start, end time.Time, page, limit int) ([]body.AbsenceFilterResponse, body.PaginationMeta, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
