package bookingroomlist

import (
	"ayam_bangkok/source/common/models"
	"ayam_bangkok/source/pkg/logger"
	"context"
)

// listBookingRoom implements Usecase.
func (u *usecaseImpl) listBookingRoom(ctx context.Context, roomID uint64, search string, page, limit int) ([]models.BookingRoomModel, int64, error) {
	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	bookingRoom, totalRows, err := u.repo.listBookingRoom(ctx, roomID, search, limit, offset)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get list booking room")
		return nil, 0, err
	}

	return bookingRoom, totalRows, nil
}