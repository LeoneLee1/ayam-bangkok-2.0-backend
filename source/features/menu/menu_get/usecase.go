package menuget

import (
	"ayam_bangkok/source/common/models"
	"context"
)

type Usecase interface {
<<<<<<< HEAD
	getMenu(ctx context.Context, day *string, week *int, page, limit int) ([]models.MenuModel, int64, error)
=======
	getMenu(ctx context.Context, page, limit int) ([]models.MenuModel, int64, error)
>>>>>>> 86e4512b758645a5366630775362ccf0bbd6f7a4
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
