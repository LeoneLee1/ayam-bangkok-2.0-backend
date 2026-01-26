package updateprofile

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// updateProfile implements Repository.
func (r *repositoryImpl) updateProfile(ctx context.Context, userID uint, data map[string]interface{}) error {

	if err := r.db.WithContext(ctx).Model(&models.UserModel{}).Where("id = ?", userID).Updates(data).Error; err != nil {
		return err
	}

	return nil
}