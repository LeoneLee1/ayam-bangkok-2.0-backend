package absenceclockinandout

import (
	"ayam_bangkok/source/common/models"
	"ayam_bangkok/source/features/absence/absence_clock_in_and_out/body"
	"context"
	"errors"
	"time"
)

// absence implements Usecase.
func (u *usecaseImpl) absence(ctx context.Context, userID uint, request body.AbsenceRequest) error {
	if userID == 0 {
		return errors.New("Invalid user id")
	}

	now := time.Now()
	currentHour := now.Hour()

	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	currentTime := now.Format("15:04:05")

	absence, err := u.repo.getAbsenceToday(ctx, userID, today)
	if err != nil {
		return err
	}

	switch {
		// case absence in
		case currentHour >= 7 && currentHour < 12:
			if absence != nil {
				return errors.New("Already absence today")
			}

			absenceIn := models.AbsenceModel{
				UserID: uint64(userID),
				Date: today,
				ClockIn: currentTime,
				LatitudeIn: request.Latitude,
				LongitudeIn: request.Longitude,
			}

			if err := u.repo.absenceIn(ctx, userID, &absenceIn); err != nil {
				return err
			}

		// case absence out
		case currentHour >= 12:
			if absence == nil {
				return errors.New("Absence in required before absence out")
			}

			if absence.ClockOut != nil {
				return errors.New("Already absence out today")
			}

			absenceOut := models.AbsenceModel{
				ClockOut: &currentTime,
				LatitudeOut: request.Latitude,
				LongitudeOut: request.Longitude,
			}

			if err := u.repo.absenceOut(ctx, userID, &absenceOut, today); err != nil {
				return err
			}

		default:
			return errors.New("Absence not allowed at this time")
	}

	return nil
}