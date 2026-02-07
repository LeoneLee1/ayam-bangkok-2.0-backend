package ordermenucancel

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

	if err := h.usecase.orderMenuCancel(ctx, uint64(userID), uint64(menuID)); err != nil {
		errMSG := "Failed to cancel order menu: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespOK(c, nil, nil, nil)
}