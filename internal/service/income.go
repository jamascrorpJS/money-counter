package service

import (
	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/internal/repository"
)

type IncomesService struct{ repository repository.Income }

func Incomes(repository *repository.Repositories) *IncomesService {
	return &IncomesService{repository: repository.Income}
}

func (c *IncomesService) Create(users_id uint, currency_id uint, summary int) int {
	return c.repository.Create(users_id, currency_id, summary)
}

func (c *IncomesService) GetByID(id string) *domain.Incomes {

	return c.repository.GetByID(id)
}

func (c *IncomesService) GetAll() *[]domain.Incomes {
	return c.repository.GetAll()
}

func (c *IncomesService) GetSelectsFields(s []string) *domain.Incomes {

	return c.repository.GetSelectsFields(s)
}

func (c *IncomesService) UpdateFields(id string, field *domain.Incomes) *domain.Incomes {
	return c.repository.UpdateFields(id, field)
}

func (c *IncomesService) Delete(id string) *domain.Incomes {
	return c.repository.Delete(id)
}
