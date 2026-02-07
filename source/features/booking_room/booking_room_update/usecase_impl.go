package bookingroomupdate

import (
	"ayam_bangkok/source/common/models"
	"ayam_bangkok/source/features/booking_room/booking_room_update/body"
	"ayam_bangkok/source/pkg/logger"
	"context"
	"fmt"
	"time"
)

// updateBookingRoom implements Usecase.
func (u *usecaseImpl) updateBookingRoom(ctx context.Context, bookingRoomID uint64, request body.BookingRoomUpdateRequest) error {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	dateParsed, err := time.ParseInLocation(
		"2006-01-02",
		request.Date,
		loc,
	)

	if err != nil {
		logger.Error().Err(err).Msg("Invalid date format")
		return fmt.Errorf("Invalid date format: %w", err)
	}

	bookingRoom := models.BookingRoomModel{
		TotalParticipants: request.TotalParticipants,
		Date: dateParsed,
		Start: request.Start,
		End: request.End,
	}

	if err := u.repo.updateBookingRoom(ctx, bookingRoomID, &bookingRoom); err != nil {
		logger.Error().Err(err).Msg("Failed to update booking room")
		return err
	}

	return nil
}