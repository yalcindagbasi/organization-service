package models

import "gorm.io/gorm"

type OrganizationMember struct {
	gorm.Model
	MemberID       uint   `json:"member_id" gorm:"not null" binding:"required"`
	OrganizationID uint   `json:"org_id" gorm:"not null" binding:"required"`
	Role           string `json:"role" gorm:"not null" binding:"required"`
}
