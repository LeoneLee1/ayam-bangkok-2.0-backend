package bookingroomupdate

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// updateBookingRoom implements Repository.
func (r *repositoryImpl) updateBookingRoom(ctx context.Context, bookingRoomID uint64, bookingRoom *models.BookingRoomModel) error {
	return r.db.WithContext(ctx).Where("id = ?", bookingRoomID).Updates(bookingRoom).Error
}