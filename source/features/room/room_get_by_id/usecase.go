package roomgetbyid

import (
	"ayam_bangkok/source/common/models"
	"context"
)

type Usecase interface {
	getRoomByID(ctx context.Context, roomID uint64) (*models.RoomModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
