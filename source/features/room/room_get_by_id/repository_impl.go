package roomgetbyid

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// getRoomByID implements Repository.
func (r *repositoryImpl) getRoomByID(ctx context.Context, roomID uint64) (*models.RoomModel, error) {
	var room models.RoomModel

	if err := r.db.WithContext(ctx).Where("id = ?", roomID).First(&room).Error; err != nil {
		return nil, err
	}

	return &room, nil
}