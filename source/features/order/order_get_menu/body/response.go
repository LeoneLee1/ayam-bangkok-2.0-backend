package body

import "ayam_bangkok/source/common/models"

type orderMenuResponse struct {
	ID   uint64 `json:"id"`
	Menu string `json:"menu"`
}

func ToResponse(data []models.MenuModel) []orderMenuResponse {
	var response []orderMenuResponse

	for _, orderMenu := range data {
		response = append(response, orderMenuResponse{
			ID: orderMenu.ID,
			Menu: orderMenu.Name,
		})
	}

	return response
}