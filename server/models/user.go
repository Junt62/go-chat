package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"type:varchar(50);not null;unique;index"`
	Password  string         `gorm:"type:varchar(255);not null"`
	Email     string         `gorm:"type:varchar(100);not null;unique;index"`
	Avatar    string         `gorm:"type:varchar(255)"`
	Nickname  string         `gorm:"type:varchar(50)"`
	Status    int            `gorm:"type:tinyint;default:1"`
	Role      string         `gorm:"type:varchar(20);default:'user'"`
	LastLogin time.Time      `gorm:"autoUpdateTime"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
	if _, err := HashPassword(u.Password); err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
