package absenceexportexcel

import (
	"ayam_bangkok/source/features/absence/absence_export_excel/body"
	"context"
	"time"
)

// getAbsenceForExport implements Repository.
func (r *repositoryImpl) getAbsenceForExport(ctx context.Context, start, end time.Time) ([]body.AbsenceExportRow, error) {
	var rows []body.AbsenceExportRow

	if err := r.db.WithContext(ctx).
		Table("absences").
		Select(`
			absences.date,
			users.name,
			absences.clock_in,
			absences.clock_out,
			absences.latitude_in,
			absences.longitude_in,
			absences.latitude_out,
			absences.longitude_out
		`).
		Joins("JOIN users ON users.id = absences.user_id").
		Where("absences.date BETWEEN ? AND ?", start, end).
		Order("absences.date ASC").
		Scan(&rows).Error; err != nil {
			return nil, err
		}

	return rows, nil
}