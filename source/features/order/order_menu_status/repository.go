package ordermenustatus

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	orderMenuStatus(ctx context.Context, userID, menuID uint64) (bool, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
