package bookingroomcreate

import (
	"ayam_bangkok/source/common/models"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	getCapacityRoom(ctx context.Context, roomID uint64) (int, error)
	checkBookingRoom(ctx context.Context, userID, roomID uint64, date, start, end string) (bool, error)
	bookingRoom(ctx context.Context, bookingRoom *models.BookingRoomModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
