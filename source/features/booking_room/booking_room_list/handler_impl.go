package bookingroomlist

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"ayam_bangkok/source/features/booking_room/booking_room_list/body"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	roomID, err := strconv.Atoi(c.Param("room_id"))
	if err != nil {
		errMSG := "Invalid room id"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

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

	search := c.Query("search")

	bookingRoom, totalRows, err := h.usecase.listBookingRoom(ctx, uint64(roomID), search, page, limit)
	if err != nil {
		errMSG := "Failed to get list booking room"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	response, meta := body.ToResponse(c, bookingRoom, page, limit, totalRows)

	httpresputils.HttpRespOK(c, gin.H{
		"booking_room_list": response,
		"meta": meta,
	}, nil, nil)
}