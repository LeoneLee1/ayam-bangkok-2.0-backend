package body

import (
	"ayam_bangkok/source/common/models"
	"time"
)

type menuResponse struct {
	ID        uint64      `json:"id"`
	Name      string      `json:"name"`
	Day       models.Days `json:"day"`
	Week      int         `json:"week"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func ToResponse(data *models.MenuModel) menuResponse {
	return menuResponse{
		ID: data.ID,
		Name: data.Name,
		Day: data.Day,
		Week: data.Week,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}