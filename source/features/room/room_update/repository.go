package roomupdate

import (
	"ayam_bangkok/source/common/models"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	getRoomImageByID(ctx context.Context, roomID uint64) (*models.RoomModel, error)
	updateRoom(ctx context.Context, roomID uint64, room *models.RoomModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
