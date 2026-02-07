package roomcreate

import (
	"ayam_bangkok/source/common/models"
	"ayam_bangkok/source/features/room/room_create/body"
	"ayam_bangkok/source/pkg/logger"
	"context"
)

// createRoom implements Usecase.
func (u *usecaseImpl) createRoom(ctx context.Context, request body.CreateRoomRequest, path string) error {
	room := models.RoomModel{
		Name: request.Name,
		Capacity: request.Capacity,
		Floor: request.Floor,
	}

	if path != "" {
		room.Image = path
	}

	if err := u.repo.createRoom(ctx, &room); err != nil {
		logger.Error().Err(err).Msg("Failed to create room")
		return err
	}

	return nil
}