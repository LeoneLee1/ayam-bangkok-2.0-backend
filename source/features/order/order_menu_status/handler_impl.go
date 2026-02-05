package ordermenustatus

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	jwtutils "ayam_bangkok/source/common/glob_utils/jwt_utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	userID, ok := jwtutils.GetCurrentUserID(c)
	if !ok {
		errMSG := "Unauthorized"
		httpresputils.HttpResponseUnAuth(c, &errMSG)
		return
	}

	menuID, err := strconv.Atoi(c.Param("menu_id"))
	if err != nil {
		msg := "Invalid menu id"
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	status, err := h.usecase.orderMenuStatus(ctx, uint64(userID), uint64(menuID))
	if err != nil {
		msg := "Failed to get order menu status"
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	httpresputils.HttpRespOK(c, gin.H{
		"status": status,
	}, nil, nil)
}