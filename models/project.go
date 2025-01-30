package models

import (
	"errors"
	"sidita-be/config"
	"time"
)

type Project struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Location  string    `gorm:"size:255;not null" json:"location"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

func GetProjects() ([]Project,error) {

	var projects []Project

	if err := config.DB.Find(&projects).Error; err != nil {
		return projects, errors.New("failed to get projects")
	}

	return projects, nil
}

func GetProjectByID(id string) (Project,error) {

	var p Project

	if err := config.DB.First(&p,id).Error; err != nil {
		return p, errors.New("project not found")
	}

	return p, nil
}