package profile

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// profile implements Repository.
func (r *repositoryImpl) profile(ctx context.Context, userID uint) (*models.UserModel, error) {
	var user models.UserModel

	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}