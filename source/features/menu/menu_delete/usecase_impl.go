package menudelete

import (
	"context"
	"errors"
)

// deleteMenu implements Usecase.
func (u *usecaseImpl) deleteMenu(ctx context.Context, menuID uint64) error {
	if menuID == 0 {
		return errors.New("Invalid menu id")
	}

	return u.repo.deleteMenu(ctx, menuID)
}