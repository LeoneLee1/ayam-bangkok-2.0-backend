package models

import (
	"time"

	"gorm.io/gorm"
)

type Days string

const (
	Monday Days = "monday"
	Tuesday Days = "tuesday"
	Wednesday Days = "wednesday"
	Thursday Days = "thursday"
	Friday Days = "friday"
)

type MenuModel struct {
	ID   uint64 `gorm:"primaryKey;" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	Day  string `gorm:"type:enum('monday','tuesday','wednesday','thursday','friday')" json:"day"`
	Week int    `json:"week"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (MenuModel) TableName() string {
	return "menus"
}