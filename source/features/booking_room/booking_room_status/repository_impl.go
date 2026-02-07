package bookingroomstatus

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// getBookingRoomStatus implements Repository.
func (r *repositoryImpl) getBookingRoomStatus(ctx context.Context, roomID uint64, date, start, end string) (bool, error) {
	var result struct{ID uint64}

	err := r.db.WithContext(ctx).
		Table("booking_rooms").
		Select("id").
		Where("room_id = ? AND date = ? AND deleted_at IS NULL AND (start < ? AND end > ?)", 
			roomID, date, end, start,
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