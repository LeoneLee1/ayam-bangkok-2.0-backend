package ordergetmenu

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// getMenuByWeekAndDays implements Repository.
func (r *repositoryImpl) getMenuByWeekAndDays(ctx context.Context, week int, days []models.Days) ([]models.MenuModel, error) {
	var menu []models.MenuModel

	if len(days) == 0 {
		return menu, nil
	}

	err := r.db.WithContext(ctx).
		Where("week = ?", week).
		Where("day IN ?", days).
		Find(&menu).Error

	return menu, err
}