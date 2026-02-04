package ordergetmenu

import (
	"ayam_bangkok/source/common/models"
	"context"
	"time"
)

type Usecase interface {
	allowedDays(today time.Weekday) []models.Days
	getMenuByWeekAndDays(ctx context.Context) ([]models.MenuModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
