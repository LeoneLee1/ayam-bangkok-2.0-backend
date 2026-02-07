package roomget

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// getRoom implements Repository.
func (r *repositoryImpl) getRoom(ctx context.Context, page, limit, offset int) ([]models.RoomModel, int64, error) {
	var (
		room []models.RoomModel
		totalRows int64
	)

	query := r.db.WithContext(ctx).Model(&models.RoomModel{})

	if err := query.Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&room).Error; err != nil {
		return nil, 0, err
	}

	return room, totalRows, nil
}