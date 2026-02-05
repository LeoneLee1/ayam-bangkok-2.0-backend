package ordermenu

import "context"

type Usecase interface {
	orderMenu(ctx context.Context, userID uint64, menuID uint64) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
