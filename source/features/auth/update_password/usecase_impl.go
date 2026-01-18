package updatepassword

import (
	hashingpassword "ayam_bangkok/source/common/glob_utils/hashing_password"
	"ayam_bangkok/source/features/auth/update_password/body"
	"context"
	"errors"
)

// updatePassword implements Usecase.
func (u *usecaseImpl) updatePassword(ctx context.Context, userID uint, request body.UpdatePasswordRequest) error {
	if userID == 0 {
		return errors.New("Invalid id")
	}

	if request.NewPassword != request.ConfirmPassword {
		return errors.New("Invalid password confirmation")
	}

	hashedPassword, err := hashingpassword.HashPassword(request.NewPassword)
	if err != nil {
		return errors.New("Failed to hash password")
	}

	if err := u.repo.updatePassword(ctx, userID, hashedPassword); err != nil {
		return err
	}

	return nil
}