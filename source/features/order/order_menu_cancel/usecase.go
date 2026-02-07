package ordermenucancel

import "context"

type Usecase interface {
	orderMenuCancel(ctx context.Context, userID, menuID uint64) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
