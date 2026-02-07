package bookingroomcancel

import (
	"ayam_bangkok/source/pkg/logger"
	"context"
)

// cancelBookingRoom implements Usecase.
func (u *usecaseImpl) cancelBookingRoom(ctx context.Context, bookingRoomID uint64) error {
	if err := u.repo.cancelBookingRoom(ctx, bookingRoomID); err != nil {
		logger.Error().Err(err).Msg("Failed to cancel booking room")
		return err
	}

	return nil
}