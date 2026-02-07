package roomcreate

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// createRoom implements Repository.
func (r *repositoryImpl) createRoom(ctx context.Context, room *models.RoomModel) error {
	return r.db.WithContext(ctx).Create(&room).Error
}