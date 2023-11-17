package repository

import (
	"fmt"

	"github.com/jamascrorpJS/money-counter/internal/domain"
	"gorm.io/gorm"
)

type RefreshTokensDB struct {
	db *gorm.DB
}

func RefreshTokensRepo(db *gorm.DB) *RefreshTokensDB {

	return &RefreshTokensDB{
		db: db,
	}
}

func (c *RefreshTokensDB) Create(users_id uint, token string, expires_at string) int {
	res := c.db.Create(&domain.RefreshTokens{

		UsersID:    users_id,
		Token:      token,
		Expires_at: expires_at,
	})

	if res.Error != nil {
		fmt.Print(res.Error)
	}
	return int(res.RowsAffected)
}

func (c *RefreshTokensDB) GetByID(id int) *domain.RefreshTokens {
	refresh_tokens := &domain.RefreshTokens{}
	err := c.db.First(refresh_tokens, id).Error
	if err != nil {
		fmt.Print(err)
	}
	return refresh_tokens
}
func (c *RefreshTokensDB) GetAll() *[]domain.RefreshTokens {
	refresh_tokens := &[]domain.RefreshTokens{}
	err := c.db.Find(refresh_tokens).Error

	if err != nil {
		fmt.Print(err)
	}
	return refresh_tokens
}
func (c *RefreshTokensDB) GetSelectsFields(s ...string) *domain.RefreshTokens {
	refresh_tokens := &domain.RefreshTokens{}
	err := c.db.Select(s).Find(refresh_tokens).Error
	if err != nil {
		fmt.Print(err)
	}
	return refresh_tokens
}
func (c *RefreshTokensDB) UpdateFields(field *domain.RefreshTokens) *domain.RefreshTokens {
	refresh_tokens := &domain.RefreshTokens{}
	err := c.db.Model(refresh_tokens).Updates(field).Error
	if err != nil {
		fmt.Print(err)
	}
	return refresh_tokens
}

func (c *RefreshTokensDB) Delete(id int) *domain.RefreshTokens {
	refresh_tokens := &domain.RefreshTokens{}
	err := c.db.Delete(refresh_tokens, id).Error
	if err != nil {
		fmt.Print(err)
	}
	return refresh_tokens
}
