package body

import "time"

type AbsenceExportRow struct {
	Date         time.Time
	Name         string
	ClockIn      string
	ClockOut     *string
	LatitudeIn   float64
	LongitudeIn  float64
	LatitudeOut  *float64
	LongitudeOut *float64
}