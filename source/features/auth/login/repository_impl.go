package login

import (
	"ayam_bangkok/source/common/models"
	"context"
	"time"
)

// getUserByNik implements Repository.
func (r *repositoryImpl) getUserByNik(ctx context.Context, nik string) (*models.UserModel, error) {
	var user models.UserModel
	
	if err := r.db.WithContext(ctx).Where("nik = ?", nik).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// login implements Repository.
func (r *repositoryImpl) login(ctx context.Context, token string, user *models.UserModel) error {

	refreshToken := models.RefreshTokenModel{
		UserID: user.ID,
		RefreshToken: token,
		ExpiredAt: time.Now().Add(30 * 24 * time.Hour),
	}

	if err := r.db.WithContext(ctx).Create(&refreshToken).Error; err != nil {
		return err
	}

	return nil
}

// revokedAllRefreshTokenByUser implements Repository.
func (r *repositoryImpl) revokedAllRefreshTokenByUser(ctx context.Context, userID uint) error {
	if err := r.db.WithContext(ctx).Model(&models.RefreshTokenModel{}).Where("user_id = ? AND revoked = false", userID).Update("revoked", true).Error; err != nil {
		return err
	}

	return nil
}