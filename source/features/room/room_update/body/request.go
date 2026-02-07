package body

import "github.com/gin-gonic/gin"

type RoomUpdateRequest struct {
	Name     string `form:"name"`
	Capacity int    `form:"capacity"`
	Floor    int    `form:"floor"`
}

func (r *RoomUpdateRequest) ToRequest(c *gin.Context) error {
	if err := c.ShouldBind(r); err != nil {
		return err
	}

	return nil
}