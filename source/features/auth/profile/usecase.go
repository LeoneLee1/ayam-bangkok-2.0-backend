package profile

import (
	"ayam_bangkok/source/common/models"
	"context"
)

type Usecase interface {
	profile(ctx context.Context, userID uint) (*models.UserModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
