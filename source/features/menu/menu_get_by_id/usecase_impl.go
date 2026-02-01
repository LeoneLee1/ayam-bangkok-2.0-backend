package menugetbyid

import (
	"ayam_bangkok/source/common/models"
	"context"
	"errors"
)

// getMenuByID implements Usecase.
func (u *usecaseImpl) getMenuByID(ctx context.Context, menuID uint) (*models.MenuModel, error) {
	if menuID == 0 {
		return nil, errors.New("Invalid id")
	}

	menu, err := u.repo.getMenuByID(ctx, menuID)
	if err != nil {
		return nil, err
	}

	return menu, nil
}