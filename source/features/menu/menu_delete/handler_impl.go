package menudelete

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
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

	if err := h.usecase.deleteMenu(ctx, uint64(menuID)); err != nil {
		errMSG := "Failed to delete menu: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespOK(c, nil, nil, nil)
}