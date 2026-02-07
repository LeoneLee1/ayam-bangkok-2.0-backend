package roomgetbyid

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"ayam_bangkok/source/features/room/room_get_by_id/body"
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

	room, err := h.usecase.getRoomByID(ctx, uint64(roomID))
	if err != nil {
		errMSG := "Failed to get room by id"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	response := body.ToResponse(c, room)

	httpresputils.HttpRespOK(c, response, nil, nil)
}