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

	user, err := u.repo.getOldPasswordByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if !hashingpassword.VerifyPassword(request.OldPassword, user.Password) {
		return errors.New("Old password is wrong")
	}

	if request.NewPassword != request.ConfirmPassword {
		return errors.New("Invalid password confirmation")
	}

	if request.NewPassword == request.OldPassword {
		return errors.New("New password cannot be same as old password")
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