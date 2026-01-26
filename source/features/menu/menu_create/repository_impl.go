package menucreate

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// createMenu implements Repository.
func (r *repositoryImpl) createMenu(ctx context.Context, menu *models.MenuModel) error {
	if err := r.db.WithContext(ctx).Create(menu).Error; err != nil {
		return err
	}

	return nil
}