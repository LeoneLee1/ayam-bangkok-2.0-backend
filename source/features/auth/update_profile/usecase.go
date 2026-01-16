package updateprofile

import (
	"ayam_bangkok/source/features/auth/update_profile/body"
	"context"
)

type Usecase interface {
	updateProfile(ctx context.Context, userID uint, request body.ProfileUpdateRequest) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
