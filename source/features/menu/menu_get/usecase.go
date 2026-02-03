package menuget

import (
	"ayam_bangkok/source/common/models"
	"context"
)

type Usecase interface {
	getMenu(ctx context.Context, day *string, week *int, page, limit int) ([]models.MenuModel, int64, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
