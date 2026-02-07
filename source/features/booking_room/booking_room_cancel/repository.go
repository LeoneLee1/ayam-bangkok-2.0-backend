package bookingroomcancel

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	cancelBookingRoom(ctx context.Context, bookingRoomID uint64) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
