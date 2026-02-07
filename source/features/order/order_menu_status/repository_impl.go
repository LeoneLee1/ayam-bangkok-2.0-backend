package ordermenustatus

import (
	"context"
)

// orderMenuStatus implements Repository.
func (r *repositoryImpl) orderMenuStatus(ctx context.Context, userID, menuID uint64) (bool, error) {
	var result struct{ID uint64}

	err := r.db.WithContext(ctx).
		Table("order_menus").
		Select("id").
		Where("user_id = ? AND menu_id = ? AND deleted_at IS NULL", userID, menuID).
		Take(&result).Error

	if err != nil {
		return false, nil
	}

	if result.ID == 0 {
		return false, nil
	}

	return true, nil
}