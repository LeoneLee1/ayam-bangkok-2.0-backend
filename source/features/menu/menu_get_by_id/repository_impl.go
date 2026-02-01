package menugetbyid

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// getMenuByID implements Repository.
func (r *repositoryImpl) getMenuByID(ctx context.Context, menuID uint) (*models.MenuModel, error) {
	var menu models.MenuModel

	if err := r.db.WithContext(ctx).Where("id = ?", menuID).First(&menu).Error; err != nil {
		return nil, err
	}

	return &menu, nil
}