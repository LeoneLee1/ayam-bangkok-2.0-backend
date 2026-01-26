package absencegetbyfilter

import (
	"ayam_bangkok/source/features/absence/absence_get_by_filter/body"
	"context"
	"errors"
	"time"
)

// getAbsenceByFilter implements Usecase.
func (u *usecaseImpl) getAbsenceByFilter(ctx context.Context, start, end time.Time, page, limit int) ([]body.AbsenceFilterResponse, body.PaginationMeta, error) {
	if page < 1 {
		page = 1
	}

	if limit <= 0 || limit > 100 {
		limit = 20
	}

	if end.Before(start) {
		return nil, body.PaginationMeta{}, errors.New("end_date must be greated than start_date")
	}

	absence, total, err := u.repo.getAbsenceByFilter(ctx, start, end, page, limit)
	if err != nil {
		return nil, body.PaginationMeta{}, err
	}

	totalPage := (total + int64(limit) - 1) / int64(limit)

	meta := body.PaginationMeta{
		Page: page,
		Limit: limit,
		Total: total,
		TotalPage: totalPage,
	}

	return absence, meta, nil
}