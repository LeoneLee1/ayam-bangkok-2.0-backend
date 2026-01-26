package body

import "github.com/gin-gonic/gin"

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Nik      string `json:"nik" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Jabatan  string `json:"jabatan" binding:"required"`
	Divisi   string `json:"divisi" binding:"required"`
}

func (r *RegisterRequest) ToRequest(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		// handle error
		return err
	}

	return nil
}