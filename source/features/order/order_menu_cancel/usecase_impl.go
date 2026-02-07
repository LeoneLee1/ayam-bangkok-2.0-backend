package ordermenucancel

import (
	"context"
	"errors"
	"time"
)

// orderMenuCancel implements Usecase.
func (u *usecaseImpl) orderMenuCancel(ctx context.Context, userID, menuID uint64) error {
	now := time.Now()
	
	if now.Hour() >= 9 {
		return errors.New("Cancel menu orders can only be placed before 9 am")
	}

	return u.repo.orderMenuCancel(ctx, userID, menuID)
}