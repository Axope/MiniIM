package models

import "gorm.io/gorm"

type GroupMember struct {
	gorm.Model
	GroupID uint `gorm:"not null; index:idx_member"`
	UserID  uint `gorm:"not null; index:idx_member"`
}

func (*GroupMember) TableName() string {
	return "group_members"
}
