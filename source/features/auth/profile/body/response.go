package body

import (
	"ayam_bangkok/source/common/models"
	"time"
)

type profileResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Nik       string `json:"nik"`
	Jabatan   string `json:"jabatan"`
	Divisi    string `json:"divisi"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToResponse(data *models.UserModel) profileResponse {
	var response profileResponse

	response = profileResponse{
		ID: uint(data.ID),
		Name: data.Name,
		Nik: data.Nik,
		Jabatan: data.Jabatan,
		Divisi: data.Divisi,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return response
}