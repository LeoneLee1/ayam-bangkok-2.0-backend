package bookingroomstatus

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	getBookingRoomStatus(ctx context.Context, roomID uint64, date, start, end string) (bool, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
