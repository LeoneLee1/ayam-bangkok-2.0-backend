package login

import (
	hashingpassword "ayam_bangkok/source/common/glob_utils/hashing_password"
	jwtutils "ayam_bangkok/source/common/glob_utils/jwt_utils"
	"ayam_bangkok/source/features/auth/login/body"
	"context"
	"errors"
)

// login implements Usecase.
func (u *usecaseImpl) login(ctx context.Context, request body.LoginRequest) (string, string, error) {
	user, err := u.repo.getUserByNik(ctx, request.Nik)
	if err != nil {
		return "", "", errors.New("Invalid nik or password")
	}
	
	if !hashingpassword.VerifyPassword(request.Password, user.Password) {
		return "", "", errors.New("Invalid nik or password")
	}

	if err := u.repo.revokedAllRefreshTokenByUser(ctx, uint(user.ID)); err != nil {
		return "", "", err
	}

	accessToken, err := jwtutils.CreateAccessToken(uint(user.ID), user.Nik, string(user.Role))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwtutils.CreateRefreshToken(uint(user.ID))
	if err != nil {
		return "", "", err
	}

	if err := u.repo.login(ctx, refreshToken, user); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}