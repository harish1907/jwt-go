package models

import "gorm.io/gorm"

type JwtUser struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
