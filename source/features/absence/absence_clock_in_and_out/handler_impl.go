package absenceclockinandout

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	jwtutils "ayam_bangkok/source/common/glob_utils/jwt_utils"
	"ayam_bangkok/source/features/absence/absence_clock_in_and_out/body"

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

	var request body.AbsenceRequest
	if err := request.ToRequest(c); err != nil {
		errMSG := "Invalid request payload: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	if err := h.usecase.absence(ctx, userID, request); err != nil {
		errMSG := "Failed to absence: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespCreated(c, nil, nil)
}