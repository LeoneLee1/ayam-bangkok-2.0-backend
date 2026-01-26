package updatepassword

import (
	"ayam_bangkok/source/common/models"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	getOldPasswordByUserID(ctx context.Context, userID uint) (*models.UserModel, error)
	updatePassword(ctx context.Context, userID uint, password string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
