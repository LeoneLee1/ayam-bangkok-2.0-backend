package models

import (
	"time"

	"gorm.io/gorm"
)

type FeedbackModel struct {
	ID     uint64 `gorm:"primaryKey;" json:"id"`
	UserID uint64 `gorm:"index" json:"user_id"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	User *UserModel `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}

func (FeedbackModel) TableName() string {
	return "feedbacks"
}