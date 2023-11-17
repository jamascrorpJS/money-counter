package domain

import (
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	ID         int            `gorm:"column:id; primary_key"`
	Created_at time.Time      `gorm:"autoCreateTime"`
	Updated_at time.Time      `gorm:"autoUpdateTime"`
	Deleted_at gorm.DeletedAt `gorm:"index"`
	Name       string         `gorm:"column:name;unique;not null"`
}
