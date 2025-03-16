package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	CreateAt time.Time
	UpdateAt time.Time
}

func FindUserByEmail(email string) *User {
	var user User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

func CreateUser(user User) error {
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = HashPassword(u.Password)
	return nil
}

func HashPassword(password string) string {
	return password
}
