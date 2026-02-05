package ordermenustatus

import (
	"ayam_bangkok/source/pkg/logger"
	"context"
	"errors"
)

// orderMenuStatus implements Usecase.
func (u *usecaseImpl) orderMenuStatus(ctx context.Context, userID, menuID uint64) (bool, error) {
	if userID == 0 {
		return false, errors.New("Invalid user id")
	}

	if menuID == 0 {
		return false, errors.New("Invalid menu id")
	}

	status, err := u.repo.orderMenuStatus(ctx, userID, menuID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get order menu status")
		return false, err
	}

	return status, nil
}