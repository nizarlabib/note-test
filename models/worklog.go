package models

import (
	"errors"
	"time"

	"sidita-be/config"
)

type Worklog struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	UserID    uint   `json:"user_id"`
	ProjectID uint   `json:"project_id"`
	WorkDate  time.Time `json:"work_date"`
	HoursWorked float64 `json:"hours_worked"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

func (w *Worklog) SaveWorklog() (*Worklog, error) {
	w.CreatedAt = time.Now()
	w.UpdatedAt = time.Now()

	if err := config.DB.Create(w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func GetWorklogs() ([]Worklog,error) {

	var worklogs []Worklog

	if err := config.DB.Find(&worklogs).Error; err != nil {
		return worklogs, errors.New("failed to get worklogs")
	}

	return worklogs, nil
}

func GetWorkLogsByUserID(uid string) ([]Worklog,error) {
	var worklogs []Worklog

	err := config.DB.Model(&Worklog{}).Where("user_id = ?", uid).Find(&worklogs).Error

	return worklogs,err
}

func GetWorkLogsByProjectID(pid uint) ([]Worklog,error) {
	var worklogs []Worklog

	err := config.DB.Model(&Worklog{}).Where("project_id = ?", pid).Find(&worklogs).Error

	return worklogs,err
}

