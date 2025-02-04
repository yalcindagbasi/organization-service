package models

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	OrganizationID uint   `json:"organization_id" gorm:"not null"`
	UserID         uint   `json:"user_id" gorm:"not null"`
	Role           string `json:"role" gorm:"not null"`
}
