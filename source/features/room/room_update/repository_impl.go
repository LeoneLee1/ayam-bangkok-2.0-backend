package roomupdate

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// getRoomImageByID implements Repository.
func (r *repositoryImpl) getRoomImageByID(ctx context.Context, roomID uint64) (*models.RoomModel, error) {
	var room models.RoomModel

	if err := r.db.WithContext(ctx).Where("id = ?", roomID).First(&room).Error; err != nil {
		return nil, err
	}

	return &room, nil
}

// updateRoom implements Repository.
func (r *repositoryImpl) updateRoom(ctx context.Context, roomID uint64, room *models.RoomModel) error {
	return r.db.WithContext(ctx).Where("id = ?", roomID).Updates(room).Error
}