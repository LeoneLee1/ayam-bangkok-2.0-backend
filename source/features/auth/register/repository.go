package register

import (
	"ayam_bangkok/source/common/models"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	registerAccount(ctx context.Context, user *models.UserModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
