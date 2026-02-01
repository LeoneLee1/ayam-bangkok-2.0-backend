package menuupdate

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"ayam_bangkok/source/features/menu/menu_update/body"
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

	var request body.MenuUpdateRequest
	if err := request.ToRequest(c); err != nil {
		errMSG := "Invalid request payload: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	if err := h.usecase.updateMenu(ctx, uint64(menuID), request); err != nil {
		errMSG := "Failed to update menu: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespAccepted(c, nil)
}