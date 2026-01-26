package updateprofile

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	updateProfile(ctx context.Context, userID uint, data map[string]interface{}) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
