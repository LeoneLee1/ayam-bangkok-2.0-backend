package register

import (
	"ayam_bangkok/source/features/auth/register/body"
	"context"
)

type Usecase interface {
	registerAccount(ctx context.Context, request body.RegisterRequest) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
