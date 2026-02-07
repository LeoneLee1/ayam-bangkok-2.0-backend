package bookingroomcreate

import (
	"ayam_bangkok/source/features/booking_room/booking_room_create/body"
	"context"
)

type Usecase interface {
	bookingRoom(ctx context.Context, userID, roomID uint64, name string, request body.BookingRoomCreateRequest) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
