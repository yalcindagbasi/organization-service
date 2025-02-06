package models

import "gorm.io/gorm"

type OrganizationMember struct {
	gorm.Model
	MemberID       uint
	OrganizationID uint
	Role           string `gorm:"not null"` // Admin, Member, Viewer gibi roller ekleyebilirsin.

	Member       Member       `gorm:"foreignKey:MemberID"`
	Organization Organization `gorm:"foreignKey:OrganizationID"`
}
