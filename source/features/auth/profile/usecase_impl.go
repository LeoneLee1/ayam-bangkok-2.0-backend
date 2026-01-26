package profile

import (
	"ayam_bangkok/source/common/models"
	"context"
	"errors"
)

// profile implements Usecase.
func (u *usecaseImpl) profile(ctx context.Context, userID uint) (*models.UserModel, error) {
	if userID == 0 {
		return nil, errors.New("Invalid id")
	}

	user, err := u.repo.profile(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, err
}