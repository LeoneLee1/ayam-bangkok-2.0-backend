package absencegetbyfilter

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"strconv"
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

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	absence, meta, err := h.usecase.getAbsenceByFilter(ctx, start, end, page, limit)
	if err != nil {
		errMSG := "Failed to get absence by filter: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespOK(c, gin.H{
		"absences": absence,
		"meta": meta,
	}, nil, nil)
}