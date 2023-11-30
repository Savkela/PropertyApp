package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Surname  string
	Email    string
	Phone    int
	Password string
}