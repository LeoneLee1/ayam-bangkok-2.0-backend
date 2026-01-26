package body

import (
	"time"
)

type AbsenceFilterResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Date time.Time `json:"date"`
	ClockIn string `json:"clock_in"`
	ClockOut *string `json:"clock_out"`
	LatitudeIn float64 `json:"latitude_in"`
	LongitudeIn float64 `json:"longitude_in"`
	LatitudeOut float64 `json:"latitude_out"`
	LongitudeOut float64 `json:"longitude_out"`
}

type PaginationMeta struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	Total int64 `json:"total"`
	TotalPage int64 `json:"total_page"`
}