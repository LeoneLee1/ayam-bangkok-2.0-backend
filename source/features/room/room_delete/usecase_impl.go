package roomdelete

import (
	"ayam_bangkok/source/pkg/logger"
	"context"
	"os"
)

// deleteRoom implements Usecase.
func (u *usecaseImpl) deleteRoom(ctx context.Context, roomID uint64) error {
	var roomImage string

	room, err := u.repo.getRoomImageByID(ctx, roomID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get room image by id")
		return err
	}

	roomImage = room.Image

	if err := u.repo.deleteRoom(ctx, roomID); err != nil {
		logger.Error().Err(err).Msg("Failed to delete room")
		return err
	}

	if roomImage != "" {
		_ = os.Remove("storage/" + roomImage)
	}

	return nil
}