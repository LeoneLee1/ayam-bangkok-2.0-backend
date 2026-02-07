package roomdelete

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

// deleteRoom implements Repository.
func (r *repositoryImpl) deleteRoom(ctx context.Context, roomID uint64) error {
	return r.db.WithContext(ctx).Where("id = ?", roomID).Delete(&models.RoomModel{}).Error
}