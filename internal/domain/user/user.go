package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Repository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}
