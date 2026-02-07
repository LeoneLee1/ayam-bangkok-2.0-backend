package bookingroomcancel

import "context"

type Usecase interface {
	cancelBookingRoom(ctx context.Context, bookingRoomID uint64) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
