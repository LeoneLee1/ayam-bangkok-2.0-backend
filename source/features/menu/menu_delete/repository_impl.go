package menudelete

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// deleteMenu implements Repository.
func (r *repositoryImpl) deleteMenu(ctx context.Context, menuID uint64) error {
	return r.db.WithContext(ctx).Where("id = ?", menuID).Delete(&models.MenuModel{}).Error
}