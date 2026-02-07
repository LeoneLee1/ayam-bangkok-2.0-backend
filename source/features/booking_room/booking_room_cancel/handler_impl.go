package bookingroomcancel

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	bookingRoomID, err := strconv.Atoi(c.Param("booking_room_id"))
	if err != nil {
		errMSG := "Invalid booking room id"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	if err := h.usecase.cancelBookingRoom(ctx, uint64(bookingRoomID)); err != nil {
		errMSG := "Failed to cancel booking room"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespOK(c, nil, nil, nil)
}