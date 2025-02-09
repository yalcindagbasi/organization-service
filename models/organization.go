package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"  binding:"required,min=2,max=50"`
	Description string `json:"description" binding:"min=2,max=511"`
}
