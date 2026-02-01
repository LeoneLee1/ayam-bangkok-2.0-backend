package menugetbyid

import (
	"ayam_bangkok/source/common/models"
	"context"
)

type Usecase interface {
	getMenuByID(ctx context.Context, menuID uint) (*models.MenuModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
