package register

import (
	"ayam_bangkok/source/common/models"
	"context"
)

// registerAccount implements Repository.
func (r *repositoryImpl) registerAccount(ctx context.Context, user *models.UserModel) error {

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}

	return nil
}