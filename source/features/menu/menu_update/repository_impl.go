package menuupdate

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// updateMenu implements Repository.
func (r *repositoryImpl) updateMenu(ctx context.Context, menuID uint64, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&models.MenuModel{}).Where("id = ?", menuID).Updates(updates).Error
}