package body

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Nik      string `json:"nik" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

func (r *LoginRequest) ToRequest(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		// handle error
		return err
	}
	
	return nil
}