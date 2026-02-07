package roomupdate

import (
	"ayam_bangkok/source/features/room/room_update/body"
	"context"
)

type Usecase interface {
	updateRoom(ctx context.Context, roomID uint64, path *string, request body.RoomUpdateRequest) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
