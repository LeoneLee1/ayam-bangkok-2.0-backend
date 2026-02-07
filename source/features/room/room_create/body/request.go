package body

import "github.com/gin-gonic/gin"

type CreateRoomRequest struct {
	Name     string `form:"name" binding:"required"`
	Capacity int    `form:"capacity" binding:"required"`
	Floor    int    `form:"floor" binding:"required"`
}

func (r *CreateRoomRequest) ToRequest(c *gin.Context) error {
	if err := c.ShouldBind(r); err != nil {
		return err
	}

	return nil
}