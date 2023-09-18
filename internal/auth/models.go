package auth

import "gorm.io/gorm"

type AdminAccount struct {
	gorm.Model
	Username string `gorm:"not null; default null; unique"`
	Password string `gorm:"not null; default null"`
	Comment  string `gorm:"not null; default null"`
	IsActive bool   `gorm:"not null; default false"`
}

type LoginRequest struct {
	Email    string
	Password string
}
