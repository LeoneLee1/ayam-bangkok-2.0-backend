package bookingroomstatus

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
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

	status, err := h.usecase.getBookingRoomStatus(ctx, uint64(roomID))
	if err != nil {
		errMSG := "Failed to get booking room status"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	var msg string

	if status == true {
		msg = "Room is used"
	} else {
		msg = "Room is available"
	}

	httpresputils.HttpRespOK(c, gin.H{
		"status": msg,
	}, nil, nil)
}