package menuupdate

import (
	"ayam_bangkok/source/features/menu/menu_update/body"
	"context"
)

type Usecase interface {
	updateMenu(ctx context.Context, menuID uint64, request body.MenuUpdateRequest) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
