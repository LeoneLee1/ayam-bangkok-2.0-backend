package roomupdate

import (
	"ayam_bangkok/source/common/models"
	"ayam_bangkok/source/features/room/room_update/body"
	"ayam_bangkok/source/pkg/logger"
	"context"
	"os"
)

// updateRoom implements Usecase.
func (u *usecaseImpl) updateRoom(ctx context.Context, roomID uint64, path *string, request body.RoomUpdateRequest) error {
	room, err := u.repo.getRoomImageByID(ctx, roomID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get room image by id")
		return err
	}

	updates := models.RoomModel{
		Name: request.Name,
		Capacity: request.Capacity,
		Floor: request.Floor,
	}

	var oldPicture string

	switch {
		case path != nil && room.Image != "":
			oldPicture = room.Image
		case path != nil:
			updates.Image = *path
	}

	if err := u.repo.updateRoom(ctx, roomID, &updates); err != nil {
		logger.Error().Err(err).Msg("Failed to update room")
		return err
	}

	if oldPicture != "" {
		_ = os.Remove("storage/" + oldPicture)
	}

	return nil
}