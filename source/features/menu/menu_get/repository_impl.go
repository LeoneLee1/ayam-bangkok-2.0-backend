package menuget

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// getMenu implements Repository.
func (r *repositoryImpl) getMenu(ctx context.Context, page, limit, offset int) ([]models.MenuModel, int64, error) {
	var (
		menu []models.MenuModel
		totalRows int64
	)

	query := r.db.WithContext(ctx).Model(&models.MenuModel{})

	if err := query.Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&menu).Error; err != nil {
		return nil, 0, err
	}

	return menu, totalRows, nil
}