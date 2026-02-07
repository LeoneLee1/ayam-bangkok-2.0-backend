package roomupdate

import (
	httpresputils "ayam_bangkok/source/common/glob_utils/http_resp_utils"
	uploadutils "ayam_bangkok/source/common/glob_utils/upload_utils"
	urlutils "ayam_bangkok/source/common/glob_utils/url_utils"
	"ayam_bangkok/source/features/room/room_update/body"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	roomID, err := strconv.Atoi(c.Param("room_id"))
	if err != nil {
		errMSG := "Invalid room id"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	var request body.RoomUpdateRequest
	if err := request.ToRequest(c); err != nil {
		errMSG := "Invalid request payload: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	var path *string

	file, err := c.FormFile("image")

	switch {
		case err == nil && file != nil:
			fileName, err := uploadutils.UploadImage(c, "image", "storage/room/images")
			if err != nil {
				errMSG := "Upload image failed"
				httpresputils.HttpRespBadRequest(c, &errMSG)
				return
			}

			p := "room/images/" + fileName
			path = &p

		case err == http.ErrMissingFile:
			path = nil

		case err != nil:
			errMSG := "Invalid file upload"
			httpresputils.HttpRespBadRequest(c, &errMSG)
			return
	}

	if err := h.usecase.updateRoom(ctx, uint64(roomID), path, request); err != nil {
		errMSG := "Failed to update room"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespOK(c, gin.H{
		"image": urlutils.BuildStorageURL(c, *path),
	}, nil, nil)
}