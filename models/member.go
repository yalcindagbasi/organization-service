package models

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null" binding:"required,min=2,max=50"`
	Email    string `json:"email" gorm:"not null" binding:"required,email"`
	Password string `json:"password" gorm:"not null" binding:"required,min=6,max=50"`
}
