package absenceclockinandout

import (
	"ayam_bangkok/source/features/absence/absence_clock_in_and_out/body"
	"context"
)

type Usecase interface {
	absence(ctx context.Context, userID uint, request body.AbsenceRequest) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
