package models

import "gorm.io/gorm"

type OrganizationMember struct {
	gorm.Model
	MemberID       uint   `json:"member_id" gorm:"not null"`
	OrganizationID uint   `json:"org_id" gorm:"not null"`
	Role           string `json:"role" gorm:"not null"` // Admin, Member, Viewer gibi roller ekleyebilirsin.

	Member       Member       `gorm:"foreignKey:MemberID"`
	Organization Organization `gorm:"foreignKey:OrganizationID"`
}
