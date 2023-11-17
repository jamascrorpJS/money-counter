package domain

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID         int            `gorm:"column:id; primary_key"`
	Created_at time.Time      `gorm:"autoCreateTime"`
	Updated_at time.Time      `gorm:"autoUpdateTime"`
	Deleted_at gorm.DeletedAt `gorm:"index"`
	Name       string         `gorm:"column:name;not null"`
	Surname    string         `gorm:"column:surname"`
	Email      string         `gorm:"column:email;unique;not null"`
	Password   string         `gorm:"column:password;not null"`
	Image      string         `gorm:"column:image"`
}
