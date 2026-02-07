package bookingroomcreate

import (
	"ayam_bangkok/source/common/models"
	"ayam_bangkok/source/features/booking_room/booking_room_create/body"
	"ayam_bangkok/source/pkg/logger"
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrRoomAlreadyBooked = errors.New("Room already booked")
)

// bookingRoom implements Usecase.
func (u *usecaseImpl) bookingRoom(ctx context.Context, userID, roomID uint64, name string, request body.BookingRoomCreateRequest) error {
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

	checkBooked, err := u.repo.checkBookingRoom(ctx, roomID, request.Date, request.Start, request.End)
	if err != nil {
		logger.Error().Err(err).Msg("Failed check booking room")
		return err
	}

	if checkBooked {
		return ErrRoomAlreadyBooked
	}

	capacity, err := u.repo.getCapacityRoom(ctx, roomID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get capacity room")
		return err
	}

	if request.TotalParticipants > capacity {
		return fmt.Errorf("Max capacity room is: %d", capacity)
	}

	bookingRoom := models.BookingRoomModel{
		UserID: userID,
		RoomID: roomID,
		Name: name,
		TotalParticipants: request.TotalParticipants,
		Date: dateParsed,
		Start: request.Start,
		End: request.End,
	}

	if err := u.repo.bookingRoom(ctx, &bookingRoom); err != nil {
		logger.Error().Err(err).Msg("Failed to create booking room")
		return err
	}

	return nil
}