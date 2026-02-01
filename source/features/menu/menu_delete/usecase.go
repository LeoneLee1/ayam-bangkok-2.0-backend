package menudelete

import "context"

type Usecase interface {
	deleteMenu(ctx context.Context, menuID uint64) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
