package models

import "time"

type RefreshTokenModel struct {
	ID           		uint64 		`gorm:"primaryKey;" json:"id"`
	UserID       		uint64 		`gorm:"index" json:"user_id"`
	RefreshToken 		string 		`gorm:"type:varchar(255);uniqueIndex" json:"refresh_token"`
	ExpiredAt    		time.Time 	`gorm:"index" json:"expired_at"`
	Revoked 			bool 		`gorm:"default:false"`

	CreatedAt 			time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 			time.Time   `gorm:"autoUpdateTime" json:"updated_at"`

	User 				*UserModel 	`gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"user,omitempty"`
}

func (RefreshTokenModel) TableName() string {
	return "refresh_tokens"
}