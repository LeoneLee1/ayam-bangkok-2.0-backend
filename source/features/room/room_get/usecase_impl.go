package roomget

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// getMenu implements Usecase.
func (u *usecaseImpl) getMenu(ctx context.Context, page, limit int) ([]models.RoomModel, int64, error) {
	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	return u.repo.getRoom(ctx, page, limit, offset)
}