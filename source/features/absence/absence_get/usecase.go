package absenceget

import (
	"ayam_bangkok/source/common/models"
	"context"
)

type Usecase interface {
	getAbsence(ctx context.Context, userID uint) ([]models.AbsenceModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
