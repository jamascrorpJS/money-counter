package domain

import "gorm.io/gorm"

type RefreshTokens struct {
	gorm.Model
	UsersID    uint
	Token      string
	Expires_at string
}
