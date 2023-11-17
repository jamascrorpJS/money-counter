package domain

import (
	"time"

	"gorm.io/gorm"
)

type Incomes struct {
	gorm.Model
	ID         int            `gorm:"column:id; primary_key"`
	Created_at time.Time      `gorm:"autoCreateTime"`
	Updated_at time.Time      `gorm:"autoUpdateTime"`
	Deleted_at gorm.DeletedAt `gorm:"index"`
	UsersID    uint           `gorm:"column:users_id"`
	Users      Users          `gorm:"foreignkey:users_id" json:"-"`
	CurrencyID uint           `gorm:"column:currency_id"`
	Currency   Currency       `gorm:"foreignkey:currency_id" json:"-"`
	Summary    int            `gorm:"column:summary"`
}
