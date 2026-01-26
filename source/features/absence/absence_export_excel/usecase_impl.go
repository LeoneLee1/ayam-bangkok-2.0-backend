package absenceexportexcel

import (
	"ayam_bangkok/source/features/absence/absence_export_excel/body"
	"context"
	"time"
)

// getAbsenceForExport implements Usecase.
func (u *usecaseImpl) getAbsenceForExport(ctx context.Context, start, end time.Time) ([]body.AbsenceExportRow, error) {
	absence, err := u.repo.getAbsenceForExport(ctx, start, end)
	if err != nil {
		return nil, err
	}

	return absence, nil
}