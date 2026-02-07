package body

import (
	urlutils "ayam_bangkok/source/common/glob_utils/url_utils"
	"ayam_bangkok/source/common/models"
	"time"

	"github.com/gin-gonic/gin"
)

type roomResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Capacity  int       `json:"capacity"`
	Floor     int       `json:"floor"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToResponse(c *gin.Context, data *models.RoomModel) roomResponse {
	response := roomResponse{
		ID: data.ID,
		Name: data.Name,
		Image: urlutils.BuildStorageURL(c, data.Image),
		Capacity: data.Capacity,
		Floor: data.Floor,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return response
}