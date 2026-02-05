package ordermenustatus

import "context"

type Usecase interface {
	orderMenuStatus(ctx context.Context, userID, menuID uint64) (bool, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
