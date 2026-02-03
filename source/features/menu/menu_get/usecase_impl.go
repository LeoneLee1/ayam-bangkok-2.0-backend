package menuget

import (
	"ayam_bangkok/source/common/models"
	"context"
<<<<<<< HEAD
	"time"
)

// getMenu implements Usecase.
func (u *usecaseImpl) getMenu(ctx context.Context, day *string, week *int, page, limit int) ([]models.MenuModel, int64, error) {
=======
)

// getMenu implements Usecase.
func (u *usecaseImpl) getMenu(ctx context.Context, page, limit int) ([]models.MenuModel, int64, error) {
>>>>>>> 86e4512b758645a5366630775362ccf0bbd6f7a4
	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

<<<<<<< HEAD
	now := time.Now()

	if day == nil {
		d := now.Weekday().String()
		day = &d
	}

	if week == nil {
		_, w := now.ISOWeek()
		week = &w
	}

	return u.repo.getMenu(ctx, day, week, page, limit, offset)
=======
	return u.repo.getMenu(ctx, page, limit, offset)
>>>>>>> 86e4512b758645a5366630775362ccf0bbd6f7a4
}