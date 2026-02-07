package bookingroomstatus

import (
	"ayam_bangkok/source/pkg/logger"
	"context"
	"time"
)

// getBookingRoomStatus implements Usecase.
func (u *usecaseImpl) getBookingRoomStatus(ctx context.Context, roomID uint64) (bool, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	now := time.Now().In(loc)

	currentDate := now.Format("2006-01-02")
	currentTime := now.Format("15:04:05")

	status, err := u.repo.getBookingRoomStatus(ctx, roomID, currentDate, currentTime, currentTime)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get booking room status")
		return false, err
	}

	return status, nil
}