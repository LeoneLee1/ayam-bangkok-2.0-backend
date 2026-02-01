package menugetbyid

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"ayam_bangkok/source/features/menu/menu_get_by_id/body"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()
	
	menuID, err := strconv.Atoi(c.Param("menu_id"))
	if err != nil {
		errMSG := "Invalid menu id: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	menu, err := h.usecase.getMenuByID(ctx, uint(menuID))
	if err != nil {
		errMSG := "Failed to get menu by id: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	response := body.ToResponse(menu)

	httpresputils.HttpRespOK(c, response, nil, nil)
}