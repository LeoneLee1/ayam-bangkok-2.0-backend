package register

import (
	hashingpassword "ayam_bangkok/source/common/glob_utils/hashing_password"
	"ayam_bangkok/source/common/models"
	"ayam_bangkok/source/features/auth/register/body"
	"context"
)

// registerAccount implements Usecase.
func (u *usecaseImpl) registerAccount(ctx context.Context, request body.RegisterRequest) error {
	hashPassword, err := hashingpassword.HashPassword(request.Password)
	if err != nil {
		return err
	}

	user := &models.UserModel{
		Name: request.Name,
		Nik: request.Nik,
		Password: hashPassword,
		Jabatan: request.Jabatan,
		Divisi: request.Divisi,
		Role: models.RoleUser,
	}

	if err := u.repo.registerAccount(ctx, user); err != nil {
		return err
	}

	return nil
}