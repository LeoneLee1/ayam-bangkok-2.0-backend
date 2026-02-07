package roomcreate

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	uploadutils "ayam_bangkok/source/common/glob_utils/upload_utils"
	urlutils "ayam_bangkok/source/common/glob_utils/url_utils"
	"ayam_bangkok/source/features/room/room_create/body"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()
	
	var request body.CreateRoomRequest
	if err := request.ToRequest(c); err != nil {
		errMSG := "Invalid request payload: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	_, err := c.FormFile("image")

	if err != nil {
		errMSG := "Image is required"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	fileName, err := uploadutils.UploadImage(c, "image", "storage/room/images")
	if err != nil {
		errMSG := "Upload image failed"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	path := "room/images/" + fileName

	if err := h.usecase.createRoom(ctx, request, path); err != nil {
		errMSG := "Failed to create room"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespCreated(c, gin.H{
		"image": urlutils.BuildStorageURL(c, path),
	}, nil)
}