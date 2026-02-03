package menuget

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"ayam_bangkok/source/features/menu/menu_get/body"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	day := c.Query("day")
	weekStr := c.Query("week")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	var dayPtr *string
	if day != "" {
		dayPtr = &day
	}

	var weekPtr *int
	if weekStr != "" {
		w, err := strconv.Atoi(weekStr)
		if err == nil && w > 0 {
			weekPtr = &w
		}
	}

	menu, totalRows, err := h.usecase.getMenu(ctx, dayPtr, weekPtr, page, limit)
	if err != nil {
		errMSG := "Failed to get menu: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	data, meta := body.ToResponse(menu, page, limit, totalRows)
	
	httpresputils.HttpRespOK(c, gin.H{
		"menus": data,
		"meta": meta,
	}, nil, nil)
}