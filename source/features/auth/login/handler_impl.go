package login

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	"ayam_bangkok/source/features/auth/login/body"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	var request body.LoginRequest
	if err := request.ToRequest(c); err != nil {
		errMSG := "Invalid request payload: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	accessToken, refreshToken, err := h.usecase.login(ctx, request)
	if err != nil {
		errMSG := "Failed to login: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	response := body.LoginResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	httpresputils.HttpRespOK(c, response, nil, nil)
}