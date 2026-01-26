package absenceget

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// getAbsence implements Repository.
func (r *repositoryImpl) getAbsence(ctx context.Context, userID uint) ([]models.AbsenceModel, error) {
	var absence []models.AbsenceModel

	if err := r.db.WithContext(ctx).
		Select(models.AbsenceModel{}.FieldDefault()).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(10).
		Find(&absence).Error; err != nil {
			return nil, err
		}

	return absence, nil
}