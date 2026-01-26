package body

import (
	"ayam_bangkok/source/common/models"
	"time"
)

type AbsenceResponse struct {
	ID   uint64 `json:"id"`
	Date time.Time `json:"date"`
	ClockIn string `json:"clock_in"`
	ClockOut string `json:"clock_out"`
}

func ToResponse(data []models.AbsenceModel) []AbsenceResponse {
	var response []AbsenceResponse

	for _, absence := range data {
		response = append(response, AbsenceResponse{
			ID: absence.ID,
			Date: absence.Date,
			ClockIn: absence.ClockIn,
			ClockOut: *absence.ClockOut,
		})
	}

	return response
}