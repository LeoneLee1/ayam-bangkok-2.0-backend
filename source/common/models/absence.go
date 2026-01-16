package models

import (
	"time"

	"gorm.io/gorm"
)

type AbsenceModel struct {
	ID        uint64 `gorm:"primaryKey;" json:"id"`
	UserID    uint64 `gorm:"index:idx_user_date,unique" json:"user_id"`
	AbsenceDate  time.Time `gorm:"type:date;index:idx_user_date,unique" json:"absence_date"`
	AbsenceIn time.Time `gorm:"type:time" json:"absence_in"`
	AbsenceOut *time.Time `gorm:"type:time" json:"absence_out"`
	AbsenceLatitudeIn float64 `gorm:"type:float" json:"absence_latitude_in"`
	AbsenceLongitudeIn float64 `gorm:"type:float" json:"absence_longitude_in"`
	AbsenceLatitudeOut *float64 `gorm:"type:float" json:"absence_latitude_out"`
	AbsenceLongitudeOut *float64 `gorm:"type:float" json:"absence_longitude_out"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	User *UserModel `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}