package menuget

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// getMenu implements Usecase.
func (u *usecaseImpl) getMenu(ctx context.Context, day *string, week *int, page, limit int) ([]models.MenuModel, int64, error) {
	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	return u.repo.getMenu(ctx, day, week, page, limit, offset)
}