package register

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"ayam_bangkok/source/features/auth/register/body"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	var request body.RegisterRequest
	if err := request.ToRequest(c); err != nil {
		errMSG := "Invalid request payload: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	if err := h.usecase.registerAccount(ctx, request); err != nil {
		errMSG := "Faile to register account: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespCreated(c, nil, nil)
}