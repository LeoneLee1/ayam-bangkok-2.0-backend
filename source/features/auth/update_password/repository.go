package updatepassword

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	updatePassword(ctx context.Context, userID uint, password string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
