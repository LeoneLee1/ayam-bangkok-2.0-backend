package profile

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	jwtutils "ayam_bangkok/source/common/glob_utils/jwt_utils"
	"ayam_bangkok/source/features/auth/profile/body"

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

	user, err := h.usecase.profile(ctx, userID)
	if err != nil {
		errMSG := "Failed to get profile: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	response := body.ToResponse(user)

	httpresputils.HttpRespOK(c, response, nil, nil)
}