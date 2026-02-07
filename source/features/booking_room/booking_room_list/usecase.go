package bookingroomlist

import (
	"ayam_bangkok/source/common/models"
	"context"
)

type Usecase interface {
	listBookingRoom(ctx context.Context, roomID uint64, search string, page, limit int) ([]models.BookingRoomModel, int64, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
