package body

import (
	"ayam_bangkok/source/common/models"
	"math"
	"time"
)

type MenuResponse struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Day       models.Days `json:"day"`
	Week      int    `json:"week"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type paginationResponse struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	TotalRows int64 `json:"total_rows"`
	TotalPages int `json:"total_pages"`
}

func ToResponse(data []models.MenuModel, page, limit int, totalRows int64) ([]MenuResponse, paginationResponse) {
	var response []MenuResponse

	for _, menu := range data {
		response = append(response, MenuResponse{
			ID: menu.ID,
			Name: menu.Name,
			Day: menu.Day,
			Week: menu.Week,
			CreatedAt: menu.CreatedAt,
			UpdatedAt: menu.UpdatedAt,
		})
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

	meta := paginationResponse{
		Page: page,
		Limit: limit,
		TotalRows: totalRows,
		TotalPages: totalPages,
	}
	
	return response, meta
}