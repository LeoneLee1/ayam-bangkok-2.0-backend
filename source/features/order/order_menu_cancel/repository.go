package ordermenucancel

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	orderMenuCancel(ctx context.Context, userID, menuID uint64) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
