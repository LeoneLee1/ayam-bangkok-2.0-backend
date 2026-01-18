package updatepassword

import (
	"ayam_bangkok/source/features/auth/update_password/body"
	"context"
)

type Usecase interface {
	updatePassword(ctx context.Context, userID uint, request body.UpdatePasswordRequest) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
