package body

import "github.com/gin-gonic/gin"

type AbsenceRequest struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (r *AbsenceRequest) ToRequest(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		// handle error
		return err
	}

	return nil
}