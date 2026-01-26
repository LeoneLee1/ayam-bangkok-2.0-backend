package absenceexportexcel

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"encoding/csv"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	startStr := c.Query("start_date")
	endStr := c.Query("end_date")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		errMSG := "Failed to parse date: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	end, err :=  time.Parse("2006-01-02", endStr)
	if err != nil {
		errMSG := "Failed to parse date: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	absence, err := h.usecase.getAbsenceForExport(ctx, start, end)
	if err != nil {
		errMSG := "Failed to get absence by filter: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	c.Header("Content-Disposition", "attachment; filename=absence_report.csv")
	c.Header("Content-Type", "text/csv")

	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	_ = writer.Write([]string{
		"Date",
		"Name",
		"Clock In",
		"Clock Out",
		"Latitude In",
		"Longitude In",
		"Latitude Out",
		"Longitude Out",
	})

	for _, excel := range absence {
		clockOut := ""
		if excel.ClockOut != nil {
			clockOut = *excel.ClockOut
		}

		latOut := ""
		lngOut := ""
		if excel.LatitudeOut != nil {
			latOut = fmt.Sprintf("%f", *excel.LatitudeOut)
		}
		if excel.LongitudeOut != nil {
			lngOut = fmt.Sprintf("%f", *excel.LongitudeOut)
		}

		_ = writer.Write([]string{
			excel.Date.Format("2006-01-02"),
			excel.Name,
			excel.ClockIn,
			clockOut,
			fmt.Sprintf("%f", excel.LatitudeIn),
			fmt.Sprintf("%f", excel.LongitudeIn),
			latOut,
			lngOut,
		})
	}

	httpresputils.HttpRespAccepted(c, nil)
}