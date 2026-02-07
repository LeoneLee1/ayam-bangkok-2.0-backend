package roomdelete

import "context"

type Usecase interface {
	deleteRoom(ctx context.Context, roomID uint64) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
