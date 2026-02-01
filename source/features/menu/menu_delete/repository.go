package menudelete

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	deleteMenu(ctx context.Context, menuID uint64) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
