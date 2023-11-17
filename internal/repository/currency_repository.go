package repository

import (
	"fmt"

	"github.com/jamascrorpJS/money-counter/internal/domain"
	"gorm.io/gorm"
)

type CurrencyDB struct {
	db *gorm.DB
}

func CurrencyRepo(db *gorm.DB) *CurrencyDB {

	return &CurrencyDB{
		db: db,
	}
}

func (c *CurrencyDB) Create(name string) int {
	res := c.db.Create(&domain.Currency{

		Name: name,
	})
	if res.Error != nil {
		fmt.Print(res.Error)
	}
	return int(res.RowsAffected)
}

func (c *CurrencyDB) GetByID(id string) *domain.Currency {
	currency := &domain.Currency{}
	err := c.db.First(currency, id).Error
	if err != nil {
		fmt.Print(err)
	}
	return currency
}
func (c *CurrencyDB) GetAll() *[]domain.Currency {
	currency := &[]domain.Currency{}
	err := c.db.Find(currency).Error
	if err != nil {
		fmt.Print(err)
	}
	return currency
}
func (c *CurrencyDB) GetSelectsFields(s []string) *domain.Currency {
	currency := &domain.Currency{}
	err := c.db.Select(s).Find(currency).Error
	if err != nil {
		fmt.Print(err)
	}
	return currency
}

func (c *CurrencyDB) UpdateFields(id string, field *domain.Currency) *domain.Currency {
	currency := &domain.Currency{}
	err := c.db.Model(currency).Where("id = ?", id).Updates(field).Error
	if err != nil {
		fmt.Print(err)
	}
	return currency
}

func (c *CurrencyDB) Delete(id string) *domain.Currency {
	currency := &domain.Currency{}
	err := c.db.Delete(currency, id).Error
	if err != nil {
		fmt.Print(err)
	}
	return currency
}
