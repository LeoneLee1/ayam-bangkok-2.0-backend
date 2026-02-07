package roomget

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"ayam_bangkok/source/features/room/room_get/body"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	room, totalRows, err := h.usecase.getMenu(ctx, page, limit)
	if err != nil {
		errMSG := "Failed to get room"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	data, meta := body.ToResponse(c, room, page, limit, totalRows)

	httpresputils.HttpRespOK(c, gin.H{
		"rooms": data,
		"meta": meta,
	}, nil, nil)
}