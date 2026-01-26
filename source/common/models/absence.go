package models

import (
	"time"

	"gorm.io/gorm"
)

type AbsenceModel struct {
	ID          uint64 `gorm:"primaryKey;" json:"id"`
	UserID      uint64 `gorm:"index" json:"user_id"`
	Date time.Time `gorm:"type:date;index" json:"date"`
	ClockIn string `gorm:"type:time" json:"clock_in"`
	ClockOut *string `gorm:"type:time" json:"clock_out"`
	LatitudeIn float64 `gorm:"type:float" json:"latitude_in"`
	LongitudeIn float64 `gorm:"type:float" json:"longitude_in"`
	LatitudeOut float64 `gorm:"type:float" json:"latitude_out"`
	LongitudeOut float64 `gorm:"type:float" json:"longitude_out"`
	
	User *UserModel `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`

	CreatedAt 	time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt 	gorm.DeletedAt `gorm:"index" json:"-"`
}

func (AbsenceModel) TableName() string {
	return "absences"
}

func (AbsenceModel) FieldDefault() []string {
	return []string{
		"id",
		"date",
		"clock_in",
		"clock_out",
	}
}