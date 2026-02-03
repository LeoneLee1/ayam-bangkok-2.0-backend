package menuget

import (
	"ayam_bangkok/source/common/models"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
<<<<<<< HEAD
	getMenu(ctx context.Context, day *string, week *int, page, limit, offset int) ([]models.MenuModel, int64, error)
=======
	getMenu(ctx context.Context, page, limit, offset int) ([]models.MenuModel, int64, error)
>>>>>>> 86e4512b758645a5366630775362ccf0bbd6f7a4
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
