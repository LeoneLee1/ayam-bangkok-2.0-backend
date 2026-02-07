package bookingroomlist

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// listBookingRoom implements Repository.
func (r *repositoryImpl) listBookingRoom(ctx context.Context, roomID uint64, search string, limit, offset int) ([]models.BookingRoomModel, int64, error) {
	var (
		bookingRoom []models.BookingRoomModel
		totalRows int64
	)

	query := r.db.WithContext(ctx).Model(&models.BookingRoomModel{})

	if err := query.Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	if search != "" {
		query = query.Where("name LIKE ?", search+"%")
	}

	if err := query.Preload("Room").Where("room_id = ?", roomID).Order("date DESC").Limit(limit).Offset(offset).Find(&bookingRoom).Error; err != nil {
		return nil, 0, err
	}

	return bookingRoom, totalRows, nil
}