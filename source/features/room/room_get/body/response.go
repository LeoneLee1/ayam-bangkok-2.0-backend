package body

import (
	urlutils "ayam_bangkok/source/common/glob_utils/url_utils"
	"ayam_bangkok/source/common/models"
	"math"
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

type paginationResponse struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
}

func ToResponse(c *gin.Context, data []models.RoomModel, page, limit int, totalRows int64) ([]roomResponse, paginationResponse) {
	var response []roomResponse

	for _, room := range data {
		response = append(response, roomResponse{
			ID: room.ID,
			Name: room.Name,
			Image: urlutils.BuildStorageURL(c, room.Image),
			Capacity: room.Capacity,
			Floor: room.Floor,
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
		})
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

	meta := paginationResponse{
		Page: page,
		Limit: limit,
		TotalRows: totalRows,
		TotalPages: totalPages,
	}

	return response, meta
}