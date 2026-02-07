package roomgetbyid

import (
	"ayam_bangkok/source/common/models"
	"ayam_bangkok/source/pkg/logger"
	"context"
)

// getRoomByID implements Usecase.
func (u *usecaseImpl) getRoomByID(ctx context.Context, roomID uint64) (*models.RoomModel, error) {
	room, err := u.repo.getRoomByID(ctx, roomID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get room by id")
		return nil, err
	}

	return room, nil
}