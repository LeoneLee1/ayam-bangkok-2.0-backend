package bookingroomupdate

import (
	"ayam_bangkok/source/common/models"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	updateBookingRoom(ctx context.Context, bookingRoomID uint64, bookingRoom *models.BookingRoomModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
