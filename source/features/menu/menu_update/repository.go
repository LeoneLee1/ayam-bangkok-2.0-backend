package menuupdate

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	updateMenu(ctx context.Context, menuID uint64, updates map[string]interface{}) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
