package models

import (
	"time"

	"gorm.io/gorm"
)

type RoomModel struct {
	ID   		uint64 		   `gorm:"primaryKey;" json:"id"`
	Name 		string 		   `gorm:"type:varchar(255)" json:"name"`
	Image  		string 		   `gorm:"type:varchar(255)" json:"image"`
	Capacity 	int    		   `json:"capacity"`
	Floor 		int    		   `json:"floor"`

	CreatedAt 	time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt 	gorm.DeletedAt `gorm:"index" json:"-"`
}

func (RoomModel) TableName() string {
	return "rooms"
}