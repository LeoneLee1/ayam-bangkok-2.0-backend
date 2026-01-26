package absenceget

import (
	"ayam_bangkok/source/common/models"
	"context"
	"errors"
)

// getAbsence implements Usecase.
func (u *usecaseImpl) getAbsence(ctx context.Context, userID uint) ([]models.AbsenceModel, error) {
	if userID == 0 {
		return nil, errors.New("Invalid user id")
	}

	absence, err := u.repo.getAbsence(ctx, userID)
	if err != nil {
		return nil, err
	}

	return absence, nil
}