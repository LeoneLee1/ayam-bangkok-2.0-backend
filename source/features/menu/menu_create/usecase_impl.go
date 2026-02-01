package menucreate

import (
	"ayam_bangkok/source/common/models"
	"ayam_bangkok/source/features/menu/menu_create/body"
	"context"
	"time"
)

// createMenu implements Usecase.
func (u *usecaseImpl) createMenu(ctx context.Context, request body.MenuCreateRequest) error {
	now := time.Now()

	_, week := now.ISOWeek()

	day := models.Days(now.Weekday().String())

	menu := models.MenuModel{
		Name: request.Name,
		Day: day,
		Week: week,
	}

	if err := u.repo.createMenu(ctx, &menu); err != nil {
		return err
	}

	return nil
}