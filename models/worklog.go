package models

import (
	"errors"
	"time"

	"sidita-be/config"
)

type Worklog struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	UserID      uint      `json:"user_id"`
	User        User   	  `gorm:"foreignKey:UserID" json:"user"` 
	ProjectID   uint      `json:"project_id"` 
	Project     Project   `gorm:"foreignKey:ProjectID" json:"project"` 
	WorkDate    string    `json:"work_date"`
	HoursWorked int       `json:"hours_worked"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"-"`
}

func (w *Worklog) SaveWorklog() (*Worklog, error) {
	w.CreatedAt = time.Now()
	w.UpdatedAt = time.Now()

	if err := config.DB.Create(w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func GetWorklogs(limit, offset int) ([]Worklog, error) {

	var worklogs []Worklog

	err := config.DB.Preload("Project").Preload("User").
		Limit(limit).                   
		Offset(offset).
		Order("work_date ASC").                 
		Find(&worklogs).Error 

	return worklogs, err
}

func GetWorklogsNotPaginated() ([]Worklog, error) {

	var worklogs []Worklog

	err := config.DB.Preload("Project").Preload("User").
		Find(&worklogs).Error 

	return worklogs, err
}

func GetWorklogByID(id string) (Worklog, error) {

	var w Worklog

	if err := config.DB.First(&w, id).Error; err != nil {
		return w, errors.New("worklog not found")
	}

	return w, nil
}

func CountAllWorklogs() (int64, error) {
	var count int64

	err := config.DB.Model(&Worklog{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func CountWorklogsUser(uid int) (int64, error) {
	var count int64

	err := config.DB.Model(&Worklog{}).Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetWorkLogsByUserID(uid int, limit, offset int) ([]Worklog, error) {
	var worklogs []Worklog

	err := config.DB.Preload("Project").Preload("User"). 
		Where("user_id = ?", uid).      
		Limit(limit).                   
		Offset(offset).
		Order("work_date DESC").                 
		Find(&worklogs).Error           

	return worklogs, err
}


func GetWorkLogsByProjectID(pid uint) ([]Worklog,error) {
	var worklogs []Worklog

	err := config.DB.Model(&Worklog{}).Where("project_id = ?", pid).Find(&worklogs).Error

	return worklogs,err
}

func DeleteWorklog(id uint) error {
	var w Worklog
	if err := config.DB.First(&w, id).Error; err != nil {
		return err
	}
	return config.DB.Delete(&w).Error
}

type WorklogSummary struct {
	ProjectName string  `json:"project_name"`
	HoursWorked float64 `json:"hours_worked"`
	TotalWorkDays float64 `json:"total_work_days"`
}
func CountUserHoursWorked(uid int) ([]WorklogSummary, error) {
	var worklogs []WorklogSummary

	err := config.DB.
		Table("worklogs w").
		Select("p.name as project_name, SUM(w.hours_worked) as hours_worked, SUM(w.hours_worked) / 8 as total_work_days").
		Joins("JOIN projects p ON w.project_id = p.id").
		Where("w.user_id = ?", uid).
		Group("p.id, p.name").
		Find(&worklogs).Error

	return worklogs, err
}


