package absenceclockinandout

import (
	"ayam_bangkok/source/common/models"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

// getAbsenceToday implements Repository.
func (r *repositoryImpl) getAbsenceToday(ctx context.Context, userID uint, today time.Time) (*models.AbsenceModel, error) {
	var absence models.AbsenceModel

	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Where("date = ?", today).First(&absence).Error; err != nil {
		switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				return nil, nil
		}
		return nil, err
	}

	return &absence, nil
}

// absenceIn implements Repository.
func (r *repositoryImpl) absenceIn(ctx context.Context, userID uint, absence *models.AbsenceModel) error {
	if err := r.db.WithContext(ctx).Create(&absence).Error; err != nil {
		return err
	}

	return nil
}

// absenceOut implements Repository.
func (r *repositoryImpl) absenceOut(ctx context.Context, userID uint, absence *models.AbsenceModel, today time.Time) error {
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Where("date = ?", today).Updates(&absence).Error; err != nil {
		return err
	}

	return nil
}