package updatepassword

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// updatePassword implements Repository.
func (r *repositoryImpl) updatePassword(ctx context.Context, userID uint, password string) error {
	if err := r.db.WithContext(ctx).Model(&models.UserModel{}).Where("id = ?", userID).Update("password", password).Error; err != nil {
		return err
	}

	return nil
}