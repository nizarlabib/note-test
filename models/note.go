package models

import (
	"errors"
	"note-test/config"
	"note-test/utils/helper"
	"time"
)

type Note struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"size:255;not null" json:"description"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func CreateNote(n *Note) (*Note, error) {
	n.CreatedAt = time.Now()
	if err := config.DB.Create(n).Error; err != nil {
		return nil, err
	}
	return n, nil
}

func GetAllNote(n *Note, pg *helper.Pagination) error {
	var notes []Note

	db := config.DB.Model(&Note{})

	if n.UserID != 0 {
		db = db.Where("user_id = ?", n.UserID)
	}

	db = db.Scopes(helper.Paginate(&notes, pg, db))

	if err := db.Order("id ASC").Find(&notes).Error; err != nil {
		return err
	}

	pg.Rows = notes
	return nil
}

func GetNoteByID(id int) (Note, error) {

	var n Note

	if err := config.DB.First(&n, id).Error; err != nil {
		return n, errors.New("note not found")
	}

	return n, nil
}

func UpdateNote(n *Note) (*Note, error) {
	n.UpdatedAt = time.Now()

	if err := config.DB.Save(n).Error; err != nil {
		return nil, err
	}

	return n, nil
}

func DeleteNote(id int) error {
	if id == 0 {
		return errors.New("invalid ID")
	}

	res := config.DB.Where("id = ?", id).Delete(&Note{})
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("note not found")
	}

	return nil
}
