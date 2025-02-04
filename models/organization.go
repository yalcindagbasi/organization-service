package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	LoginName   string `json:"login_name" gorm:"not null"`
	Password    string `json:"password" gorm:"not null"`
}
