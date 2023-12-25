package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Uuid   string `gorm:"type:varchar(150); not null; unique_index:idx_uuid;"`
	UserID uint   `gorm:"not null"`
	Name   string `gorm:"type:varchar(150); not null;"`
}

func (u *Group) TableName() string {
	return "groups"
}
