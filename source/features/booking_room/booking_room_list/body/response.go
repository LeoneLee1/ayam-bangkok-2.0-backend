package body

import (
	urlutils "ayam_bangkok/source/common/glob_utils/url_utils"
	"ayam_bangkok/source/common/models"
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

type bookingRoomList struct {
	ID                uint64 `json:"id"`
	Name              string `json:"name"`
	TotalParticipants int    `json:"total_participants"`
	Date              string `json:"date"`
	Start             string `json:"start"`
	End               string `json:"end"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Room              roomResponse `json:"rooms"`
}

type roomResponse struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	Floor     int    `json:"floor"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type paginationResponse struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	TotalRows int64 `json:"total_rows"`
	TotalPages int `json:"total_pages"`
}

func ToResponse(c *gin.Context, data []models.BookingRoomModel, page, limit int, totalRows int64) ([]bookingRoomList, paginationResponse) {
	var response []bookingRoomList

	for _, bookingRoom := range data {

		var room roomResponse

		room = roomResponse{
			ID: bookingRoom.Room.ID,
			Name: bookingRoom.Room.Name,
			Image: urlutils.BuildStorageURL(c, bookingRoom.Room.Image),
			Floor: bookingRoom.Room.Floor,
			CreatedAt: bookingRoom.Room.CreatedAt,
			UpdatedAt: bookingRoom.Room.UpdatedAt,
		}

		response = append(response, bookingRoomList{
			ID: bookingRoom.ID,
			Name: bookingRoom.Name,
			TotalParticipants: bookingRoom.TotalParticipants,
			Date: bookingRoom.Date.Format("2006-01-02"),
			Start: bookingRoom.Start,
			End: bookingRoom.End,
			CreatedAt: bookingRoom.CreatedAt,
			UpdatedAt: bookingRoom.UpdatedAt,
			Room: room,
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