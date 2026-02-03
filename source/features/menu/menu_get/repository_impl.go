package menuget

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// getMenu implements Repository.
<<<<<<< HEAD
func (r *repositoryImpl) getMenu(ctx context.Context, day *string, week *int, page, limit, offset int) ([]models.MenuModel, int64, error) {
=======
func (r *repositoryImpl) getMenu(ctx context.Context, page, limit, offset int) ([]models.MenuModel, int64, error) {
>>>>>>> 86e4512b758645a5366630775362ccf0bbd6f7a4
	var (
		menu []models.MenuModel
		totalRows int64
	)

	query := r.db.WithContext(ctx).Model(&models.MenuModel{})

<<<<<<< HEAD
	if day != nil && *day != "" {
		query = query.Where("day = ?", *day)
	}

	if week != nil && *week > 0 {
		query = query.Where("week = ?", *week)
	}

=======
>>>>>>> 86e4512b758645a5366630775362ccf0bbd6f7a4
	if err := query.Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&menu).Error; err != nil {
		return nil, 0, err
	}

	return menu, totalRows, nil
}