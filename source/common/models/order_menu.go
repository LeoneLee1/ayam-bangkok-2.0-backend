package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderMenuModel struct {
	ID        uint64         `gorm:"primaryKey;" json:"id"`
	UserID    uint64         `gorm:"index" json:"user_id"`
	MenuID    uint64         `gorm:"index" json:"menu_id"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	User *UserModel `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Menu *MenuModel `gorm:"foreignKey:MenuID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"menu,omitempty"`
}

func (OrderMenuModel) TableName() string {
	return "order_menus"
}