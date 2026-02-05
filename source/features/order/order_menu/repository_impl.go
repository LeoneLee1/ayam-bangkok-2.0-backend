package ordermenu

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// orderMenu implements Repository.
func (r *repositoryImpl) orderMenu(ctx context.Context, userID, menuID uint64) error {
	orderMenu := models.OrderMenuModel{
		UserID: userID,
		MenuID: menuID,
	}

	return r.db.WithContext(ctx).Create(&orderMenu).Error
}