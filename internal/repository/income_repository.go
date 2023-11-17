package repository

import (
	"fmt"

	"github.com/jamascrorpJS/money-counter/internal/domain"
	"gorm.io/gorm"
)

type IncomeDB struct {
	db *gorm.DB
}

func IncomeRepo(db *gorm.DB) *IncomeDB {

	return &IncomeDB{
		db: db,
	}
}

func (c *IncomeDB) Create(users_id uint, currency_id uint, summary int) int {
	res := c.db.Create(&domain.Incomes{
		CurrencyID: currency_id,
		UsersID:    users_id,
		Summary:    summary,
	})

	if res.Error != nil {
		fmt.Print(res.Error)
	}
	return int(res.RowsAffected)
}

func (c *IncomeDB) GetByID(id string) *domain.Incomes {
	incomes := &domain.Incomes{}
	err := c.db.First(incomes, id).Error
	if err != nil {
		fmt.Print(err)
	}
	return incomes
}
func (c *IncomeDB) GetAll() *[]domain.Incomes {
	incomes := &[]domain.Incomes{}
	err := c.db.Find(incomes).Error
	if err != nil {

		fmt.Print(err)
	}
	return incomes
}
func (c *IncomeDB) GetSelectsFields(s []string) *domain.Incomes {
	incomes := &domain.Incomes{}
	err := c.db.Select(s).Find(incomes).Error
	if err != nil {
		fmt.Print(err)
	}
	return incomes
}
func (c *IncomeDB) UpdateFields(id string, field *domain.Incomes) *domain.Incomes {
	incomes := &domain.Incomes{}
	err := c.db.Model(incomes).Where("id = ?", id).Updates(field).Error
	if err != nil {
		fmt.Print(err)
	}
	return incomes
}

func (c *IncomeDB) Delete(id string) *domain.Incomes {
	incomes := &domain.Incomes{}
	err := c.db.Delete(incomes, id).Error
	if err != nil {
		fmt.Print(err)
	}
	return incomes
}
