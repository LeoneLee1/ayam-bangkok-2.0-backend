package models

import (
	"time"

	"gorm.io/gorm"
)

type RoleAccess string

const (
	RoleUser  RoleAccess = "user"
	RoleAdmin RoleAccess = "admin"
)

type UserModel struct {
	ID        uint64         `gorm:"primaryKey;" json:"id"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	Nik       string         `gorm:"type:varchar(255);uniqueIndex" json:"nik"`
	Password  string         `gorm:"type:varchar(255)" json:"-"`
	Jabatan   string         `gorm:"type:varchar(255)" json:"jabatan"`
	Divisi    string         `gorm:"type:varchar(255)" json:"divisi"`
	Role      RoleAccess     `gorm:"type:enum('admin','user');index" json:"role"`
	
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Absences []AbsenceModel  `gorm:"foreignKey:UserID;references:ID" json:"absences,omitempty"`
	Feedback []FeedbackModel `gorm:"foreignKey:UserID;references:ID" json:"feedbacks,omitempty"`
	OrderMenu []OrderMenuModel  `gorm:"foreignKey:UserID;references:ID" json:"order_menus,omitempty"`
	BookingRoom []BookingRoomModel `gorm:"foreignKey:UserID;references:ID" json:"booking_room,omitempty"`
}

func (UserModel) TableName() string {
	return "users"
}