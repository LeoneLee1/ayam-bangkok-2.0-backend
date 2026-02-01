package menuupdate

import (
	"ayam_bangkok/source/features/menu/menu_update/body"
	"context"
	"errors"
)

// updateMenu implements Usecase.
func (u *usecaseImpl) updateMenu(ctx context.Context, menuID uint64, request body.MenuUpdateRequest) error {
	if menuID == 0 {
		return errors.New("Invalid menu id")
	}

	updates := map[string]interface{}{
		"name": request.Name,
	}

	return u.repo.updateMenu(ctx, menuID, updates)
}