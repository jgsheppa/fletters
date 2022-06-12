package models

import "gorm.io/gorm"

// User model which stores user name, email address,
// password hash, and remember hash in the PSQL database.
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique"`
	Role         string `gorm:"not null"`
}
