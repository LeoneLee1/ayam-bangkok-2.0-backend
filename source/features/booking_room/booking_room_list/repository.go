package bookingroomlist

import (
	"ayam_bangkok/source/common/models"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	listBookingRoom(ctx context.Context, roomID uint64, search string, limit, offset int) ([]models.BookingRoomModel, int64, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
