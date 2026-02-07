package bookingroomupdate

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"ayam_bangkok/source/features/booking_room/booking_room_update/body"
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

	var request body.BookingRoomUpdateRequest
	if err := request.ToRequest(c); err != nil {
		errMSG := "Invalid request payload: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	if err := h.usecase.updateBookingRoom(ctx, uint64(bookingRoomID), request); err != nil {
		errMSG := "Failed to update booking room"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespAccepted(c, nil)
}