package ordermenucancel

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// orderMenuCancel implements Repository.
func (r *repositoryImpl) orderMenuCancel(ctx context.Context, userID, menuID uint64) error {
	if err := r.db.WithContext(ctx).Where("user_id = ? AND menu_id = ?", userID, menuID).Delete(&models.OrderMenuModel{}).Error; err != nil {
		return err
	}

	return nil
}