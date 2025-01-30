package models

import (
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"sidita-be/config"
	"sidita-be/utils/token"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

func LoginCheck(email string, password string) (string,error) {
	
	var err error

	u := User{}

	err = config.DB.Model(User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token,err := token.GenerateToken(u.ID)

	if err != nil {
		return "",err
	}

	return token,nil
	
}

func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) SaveUser() (*User, error) {
	var existingUser User
	err := config.DB.Model(&User{}).Where("email = ?", u.Email).Take(&existingUser).Error
	if err == nil {
		return nil, fmt.Errorf("email %s is already registered", u.Email)
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	if err := config.DB.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func GetUserByID(uid int) (User,error) {

	var u User

	if err := config.DB.First(&u,uid).Error; err != nil {
		return u,errors.New("User not found")
	}

	return u,nil
}

func GetUsers(limit, offset int) ([]User,error) {

	var users []User

	if err := config.DB.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return users, errors.New("failed to get users")
	}

	return users, nil
}

func CountAllUsers() (int64, error) {
	var count int64

	err := config.DB.Model(&User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}