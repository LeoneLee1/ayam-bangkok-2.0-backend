package bookingroomupdate

import (
	"ayam_bangkok/source/features/booking_room/booking_room_update/body"
	"context"
)

type Usecase interface {
	updateBookingRoom(ctx context.Context, bookingRoomID uint64, request body.BookingRoomUpdateRequest) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
