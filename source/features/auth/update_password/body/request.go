package body

import "github.com/gin-gonic/gin"

type UpdatePasswordRequest struct {
	OldPassword 	string `json:"old_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

func (r *UpdatePasswordRequest) ToRequest(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		// handle error
		return err
	}

	return nil
}