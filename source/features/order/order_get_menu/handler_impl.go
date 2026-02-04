package ordergetmenu

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"ayam_bangkok/source/features/order/order_get_menu/body"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	menu, err := h.usecase.getMenuByWeekAndDays(ctx)
	if err != nil {
		errMSG := "Failed to get order menu data: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	response := body.ToResponse(menu)

	httpresputils.HttpRespOK(c, response, nil, nil)
}