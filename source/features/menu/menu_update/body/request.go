package body

import "github.com/gin-gonic/gin"

type MenuUpdateRequest struct {
	Name string `json:"name"`
}

func (r *MenuUpdateRequest) ToRequest(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}

	return nil
}