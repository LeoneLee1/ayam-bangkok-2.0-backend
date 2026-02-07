package roomcreate

import (
	"ayam_bangkok/source/features/room/room_create/body"
	"context"
)

type Usecase interface {
	createRoom(ctx context.Context, request body.CreateRoomRequest, path string) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
