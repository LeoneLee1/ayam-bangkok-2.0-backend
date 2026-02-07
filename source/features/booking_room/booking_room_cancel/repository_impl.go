package bookingroomcancel

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// cancelBookingRoom implements Repository.
func (r *repositoryImpl) cancelBookingRoom(ctx context.Context, bookingRoomID uint64) error {
	return r.db.WithContext(ctx).Where("id = ?", bookingRoomID).Delete(&models.BookingRoomModel{}).Error
}