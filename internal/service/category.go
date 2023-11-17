package service

import (
	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/internal/repository"
)

type CategoryService struct{ repository repository.Category }

func Category(repository *repository.Repositories) *CategoryService {
	return &CategoryService{repository: repository.Category}
}

func (c *CategoryService) Create(name string) int {
	return c.repository.Create(name)
}

func (c *CategoryService) GetByID(id string) *domain.Category {

	return c.repository.GetByID(id)
}

func (c *CategoryService) GetAll() *[]domain.Category {
	return c.repository.GetAll()
}

func (c *CategoryService) GetSelectsFields(s []string) *domain.Category {

	return c.repository.GetSelectsFields(s)
}

func (c *CategoryService) UpdateFields(id string, field *domain.Category) *domain.Category {
	return c.repository.UpdateFields(id, field)
}

func (c *CategoryService) Delete(id string) *domain.Category {
	return c.repository.Delete(id)
}
