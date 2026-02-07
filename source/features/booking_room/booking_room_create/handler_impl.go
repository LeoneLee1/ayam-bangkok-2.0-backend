package bookingroomcreate

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	jwtutils "ayam_bangkok/source/common/glob_utils/jwt_utils"
	"ayam_bangkok/source/features/booking_room/booking_room_create/body"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()
	
	userClaim, ok := jwtutils.GetCurrentUser(c)
	if !ok {
		errMSG := "Unauthorized"
		httpresputils.HttpResponseUnAuth(c, &errMSG)
		return
	}

	roomID, err := strconv.Atoi(c.Param("room_id"))
	if err != nil {
		errMSG := "Invalid room id"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	var request body.BookingRoomCreateRequest
	if err := request.ToRequest(c); err != nil {
		errMSG := "Invalid request payload: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	if err := h.usecase.bookingRoom(ctx, uint64(userClaim.ID), uint64(roomID), userClaim.Name, request); err != nil {
		switch {
			case errors.Is(err, ErrRoomAlreadyBooked):
				errMSG := "Room already booked in this time range"
				httpresputils.HttpRespBadRequest(c, &errMSG)
				return
			default:
				errMSG := err.Error()
				httpresputils.HttpRespBadRequest(c, &errMSG)
				return
		}
	}

	httpresputils.HttpRespCreated(c, nil, nil)
}