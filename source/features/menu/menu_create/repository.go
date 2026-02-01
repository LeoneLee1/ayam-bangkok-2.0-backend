package menucreate

import (
	"ayam_bangkok/source/common/models"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	createMenu(ctx context.Context, menu *models.MenuModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
