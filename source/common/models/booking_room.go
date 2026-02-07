package models

import (
	"time"

	"gorm.io/gorm"
)

type BookingRoomModel struct {
	ID     uint64 `gorm:"primaryKey;" json:"id"`
	UserID uint64 `gorm:"index" json:"user_id"`
	RoomID uint64 `gorm:"index" json:"room_id"`

	Name string `gorm:"type:varchar(255);index" json:"name"`
	TotalParticipants int `json:"total_participants"`
	Date time.Time `gorm:"type:date;index" json:"date"`
	Start string `gorm:"type:time;index" json:"start"`
	End string `gorm:"type:time;index" json:"end"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	User *UserModel `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Room *RoomModel `gorm:"foreignKey:RoomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"room,omitempty"`
}

func (BookingRoomModel) TableName() string {
	return "booking_rooms"
}