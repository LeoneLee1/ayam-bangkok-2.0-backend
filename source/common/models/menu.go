package models

import (
	"time"

	"gorm.io/gorm"
)

type Days string

const (
	Monday Days = "Monday"
	Tuesday Days = "Tuesday"
	Wednesday Days = "Wednesday"
	Thursday Days = "Thursday"
	Friday Days = "Friday"
)

type MenuModel struct {
	ID   uint64 `gorm:"primaryKey;" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	Day  Days `gorm:"type:varchar(20)" json:"day"`
	Week int    `gorm:"index" json:"week"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	OrderMenu []OrderMenuModel  `gorm:"foreignKey:MenuID;references:ID" json:"order_menus,omitempty"`
}

func (MenuModel) TableName() string {
	return "menus"
}