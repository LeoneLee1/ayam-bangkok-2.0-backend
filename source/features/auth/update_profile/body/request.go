package body

import "github.com/gin-gonic/gin"

type ProfileUpdateRequest struct {
	Name    string `json:"name"`
	Jabatan string `json:"jabatan"`
	Divisi  string `json:"divisi"`
}

func (r *ProfileUpdateRequest) ToRequest(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		// handle error
		return err
	}

	return nil
}