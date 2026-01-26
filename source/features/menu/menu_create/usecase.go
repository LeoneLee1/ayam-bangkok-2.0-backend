package menucreate

import (
	"ayam_bangkok/source/features/menu/menu_create/body"
	"context"
)

type Usecase interface {
	createMenu(ctx context.Context, request body.MenuCreateRequest) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
