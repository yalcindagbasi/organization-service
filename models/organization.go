package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
}
