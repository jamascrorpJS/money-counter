package service

import (
	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/internal/repository"
)

type CostsService struct{ repository repository.Cost }

func Costs(repository *repository.Repositories) *CostsService {
	return &CostsService{repository: repository.Cost}
}

func (c *CostsService) Create(user_id uint, category_id uint, currency_id uint, summary int) int {
	return c.repository.Create(user_id, category_id, currency_id, summary)
}

func (c *CostsService) GetByID(id string) *domain.Costs {

	return c.repository.GetByID(id)
}

func (c *CostsService) GetAll() *[]domain.Costs {
	return c.repository.GetAll()
}

func (c *CostsService) GetSelectsFields(s []string) *domain.Costs {

	return c.repository.GetSelectsFields(s)
}

func (c *CostsService) UpdateFields(id string, field *domain.Costs) *domain.Costs {
	return c.repository.UpdateFields(id, field)
}

func (c *CostsService) Delete(id string) *domain.Costs {
	return c.repository.Delete(id)
}

func (c *CostsService) Report() error {
	return c.repository.Report()
}
