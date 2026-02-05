package ordermenu

import (
	"context"
	"errors"
	"time"
)

// orderMenu implements Usecase.
func (u *usecaseImpl) orderMenu(ctx context.Context, userID, menuID uint64) error {
	if userID == 0 {
		return errors.New("Invalid user id")
	}

	if menuID == 0 {
		return errors.New("Invalid menu id")
	}

	now := time.Now()

	if now.Hour() >= 9 {
		return errors.New("Menu orders can only be placed before 9 am")
	}

	return u.repo.orderMenu(ctx, userID, menuID)
}