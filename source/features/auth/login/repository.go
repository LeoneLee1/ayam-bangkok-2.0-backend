package login

import (
	"ayam_bangkok/source/common/models"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	login(ctx context.Context, refreshToken string, user *models.UserModel) error
	getUserByNik(ctx context.Context, nik string) (*models.UserModel, error)
	revokedAllRefreshTokenByUser(ctx context.Context, userID uint) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
