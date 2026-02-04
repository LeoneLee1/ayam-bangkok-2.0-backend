package ordergetmenu

import (
	"ayam_bangkok/source/common/models"
	"ayam_bangkok/source/pkg/logger"
	"context"
	"time"
)

// allowedDays implements Usecase.
func (u *usecaseImpl) allowedDays(today time.Weekday) []models.Days {
	switch today {
	case time.Monday:
		return []models.Days{models.Monday, models.Tuesday, models.Wednesday, models.Thursday, models.Friday}
	case time.Tuesday:
		return []models.Days{models.Tuesday, models.Wednesday, models.Thursday, models.Friday}
	case time.Wednesday:
		return []models.Days{models.Wednesday, models.Thursday, models.Friday}
	case time.Thursday:
		return []models.Days{models.Thursday, models.Friday}
	case time.Friday:
		return []models.Days{models.Friday}
	default:
		return []models.Days{}
	}
}

// getMenuByWeekAndDays implements Usecase.
func (u *usecaseImpl) getMenuByWeekAndDays(ctx context.Context) ([]models.MenuModel, error) {
	now := time.Now()
	today := now.Weekday()

	if today == time.Saturday || today == time.Sunday {
		return []models.MenuModel{}, nil
	}

	_, week := now.ISOWeek()
	days := u.allowedDays(today)

	menu, err := u.repo.getMenuByWeekAndDays(ctx, week, days)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get order menu data")
		return nil, err
	}

	return menu, nil
}