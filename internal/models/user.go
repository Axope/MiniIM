package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid     string `gorm:"type:varchar(150); not null; unique_index:idx_uuid;"`
	Username string `gorm:"type:varchar(150); not null; unique;" json:"username"`

	Password string `gorm:"type:varchar(150); not null;" json:"password"`
	// Nickname string `gorm:"type:varchar(150);"`
	// Email    string `gorm:"type:varchar(50);"`
	// Phone    string `gorm:"type:varchar(50);"`
}

func (u *User) TableName() string {
	return "users"
}
