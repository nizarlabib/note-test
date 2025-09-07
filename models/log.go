package models

import (
	"time"
	"sidita-be/config"
)

type Log struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	EndPoint string    `gorm:"size:255;not null" json:"end_point"`
	Method string    `gorm:"size:255;not null" json:"method"`
	DateTime time.Time `gorm:"autoCreateTime" json:"date_time"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User 	  User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func CreateLog(l *Log) error {
	l.DateTime = time.Now()
	if err := config.DB.Create(l).Error; err != nil {
		return nil
	}
	return nil
}