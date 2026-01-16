package updateprofile

import (
	"ayam_bangkok/source/features/auth/update_profile/body"
	"context"
	"errors"
)

// updateProfile implements Usecase.
func (u *usecaseImpl) updateProfile(ctx context.Context, userID uint, request body.ProfileUpdateRequest) error {
	if userID == 0 {
		return errors.New("Invalid id")
	}

	updates := map[string]interface{}{
		"name": request.Name,
		"jabatan": request.Jabatan,
		"divisi": request.Divisi,
	}

	if err := u.repo.updateProfile(ctx, userID, updates); err != nil {
		return err
	}

	return nil
}