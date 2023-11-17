package service

import (
	"errors"

	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/internal/repository"

	jwt "github.com/jamascrorpJS/money-counter/pkg/token"
)

type UsersService struct{ repository repository.Users }

func Users(repository *repository.Repositories) *UsersService {
	return &UsersService{repository: repository.Users}
}

func (c *UsersService) Create(name string, surname string, email string, password string, image string) (int, error) {
	return c.repository.Create(name, surname, email, password, image)
}

func (c *UsersService) GetByID(id string) *domain.Users {

	return c.repository.GetByID(id)
}

func (c *UsersService) GetAll() *[]domain.Users {
	return c.repository.GetAll()
}

func (c *UsersService) GetSelectsFields(s ...string) *domain.Users {

	return c.repository.GetSelectsFields(s)
}

func (c *UsersService) UpdateFields(id string, field *domain.Users) *domain.Users {
	return c.repository.UpdateFields(id, field)
}

func (c *UsersService) Delete(id string) *domain.Users {
	return c.repository.Delete(id)
}

func (c *UsersService) ExistedUsers(email string) (*domain.Users, error) {
	return c.repository.ExistedUsers(email)
}
func (c *UsersService) LoginUsers(email string, password string) (string, error) {

	user, err := c.repository.ExistedUsers(email)
	if err != nil || user.Password != password {

		return "", errors.New(err.Error())
	}
	jw := jwt.Jw(user.ID)
	return jw, nil
}
