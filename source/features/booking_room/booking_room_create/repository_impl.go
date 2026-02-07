package bookingroomcreate

import (
	"ayam_bangkok/source/common/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

// getCapacityRoom implements Repository.
func (r *repositoryImpl) getCapacityRoom(ctx context.Context, roomID uint64) (int, error) {
	var capacity struct{Capacity int}

	if err := r.db.WithContext(ctx).Table("rooms").Select("capacity").Where("id = ?", roomID).Take(&capacity).Error; err != nil {
		return 0, err
	}

	return capacity.Capacity, nil
}

// checkBookingRoom implements Repository.
func (r *repositoryImpl) checkBookingRoom(ctx context.Context, userID, roomID uint64, date, start, end string) (bool, error) {
	var result struct{ID uint64}

	err := r.db.WithContext(ctx).
		Table("booking_rooms").
		Select("id").
		Where("user_id = ? AND room_id = ? AND date = ? AND deleted_at IS NULL AND (start < ? AND end > ?)", 
			userID, roomID, date, end, start,
		).
		Take(&result).Error

	if err != nil {
		
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}
	
	return true, nil
}

// bookingRoom implements Repository.
func (r *repositoryImpl) bookingRoom(ctx context.Context, bookingRoom *models.BookingRoomModel) error {
	return r.db.WithContext(ctx).Create(bookingRoom).Error
}