package body

import "github.com/gin-gonic/gin"

type MenuCreateRequest struct {
	Name     string `json:"name" binding:"required"`
}

func (r *MenuCreateRequest) ToRequest(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}

	return nil
}