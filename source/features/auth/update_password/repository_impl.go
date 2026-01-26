package updatepassword

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// getOldPasswordByUserID implements Repository.
func (r *repositoryImpl) getOldPasswordByUserID(ctx context.Context, userID uint) (*models.UserModel, error) {
	var user models.UserModel
	if err := r.db.WithContext(ctx).Select("id", "password").Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// updatePassword implements Repository.
func (r *repositoryImpl) updatePassword(ctx context.Context, userID uint, password string) error {
	if err := r.db.WithContext(ctx).Model(&models.UserModel{}).Where("id = ?", userID).Update("password", password).Error; err != nil {
		return err
	}

	return nil
}