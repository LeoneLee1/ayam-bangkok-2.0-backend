package absencegetbyfilter

import (
	"ayam_bangkok/source/features/absence/absence_get_by_filter/body"
	"context"
	"time"
)

// getAbsenceByFilter implements Repository.
func (r *repositoryImpl) getAbsenceByFilter(ctx context.Context, start, end time.Time, page, limit int) ([]body.AbsenceFilterResponse, int64, error) {
	var (
		absence []body.AbsenceFilterResponse
		total int64
	)
	
	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).
		Table("absences").
		Select(`
			absences.id,
			absences.user_id,
			users.name,
			absences.date,
			absences.clock_in,
			absences.clock_out
		`).
		Joins("JOIN users ON users.id = absences.user_id").
		Where("absences.date BETWEEN ? AND ?", start, end).
		Order("absences.date ASC").
		Limit(limit).
		Offset(offset).
		Scan(&absence).Error; err != nil {
			return nil, 0, err
		}

	return absence, total, nil
}