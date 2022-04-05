package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id         uint      `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	UserName   string    `json:"user_name" gorm:"unique"`
	Email      string    `json:"email" gorm:"unique"`
	Password   []byte    `json:"-"`
	Country    string    `json:"country"`
	ProfilePic string    `json:"profile_pic"`
	Phone      string    `json:"phone" gorm:"unique"`
	Date       time.Time `gorm:"type:time" json:"date"`
	Token      float64   `json:"-"`
	Skill string `json:"skill"`
	AboutMe string `json:"about_me"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
