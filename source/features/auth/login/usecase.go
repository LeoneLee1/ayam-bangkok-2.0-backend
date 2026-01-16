package login

import (
	"ayam_bangkok/source/features/auth/login/body"
	"context"
)

type Usecase interface {
	login(ctx context.Context, request body.LoginRequest) (string, string, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
