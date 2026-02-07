package body

import (
	"github.com/gin-gonic/gin"
)

type BookingRoomCreateRequest struct {
	TotalParticipants int    `json:"total_participants"`
	Date              string `json:"date"`
	Start             string `json:"start"`
	End               string `json:"end"`
}

func (r *BookingRoomCreateRequest) ToRequest(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}

	return nil
}