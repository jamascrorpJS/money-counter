package service

import (
	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/internal/repository"
)

type CurrencyService struct{ repository repository.Currency }

func Currency(repository *repository.Repositories) *CurrencyService {
	return &CurrencyService{repository: repository.Currency}
}

func (c *CurrencyService) Create(name string) int {
	return c.repository.Create(name)
}

func (c *CurrencyService) GetByID(id string) *domain.Currency {

	return c.repository.GetByID(id)
}

func (c *CurrencyService) GetAll() *[]domain.Currency {
	return c.repository.GetAll()
}

func (c *CurrencyService) GetSelectsFields(s []string) *domain.Currency {

	return c.repository.GetSelectsFields(s)
}

func (c *CurrencyService) UpdateFields(id string, field *domain.Currency) *domain.Currency {
	return c.repository.UpdateFields(id, field)
}

func (c *CurrencyService) Delete(id string) *domain.Currency {
	return c.repository.Delete(id)
}
