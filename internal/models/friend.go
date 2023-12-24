package models

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	UserID   uint `gorm:"not null; index:idx_friend"`
	FriendID uint `gorm:"not null; index:idx_friend"`
}

func (u *Friend) TableName() string {
	return "friends"
}
