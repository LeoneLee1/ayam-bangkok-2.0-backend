package menuget

import (
	"ayam_bangkok/source/common/models"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	getMenu(ctx context.Context, page, limit, offset int) ([]models.MenuModel, int64, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
