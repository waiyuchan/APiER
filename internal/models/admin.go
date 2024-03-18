package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string
	Role     string
}
