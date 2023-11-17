package repository

import (
	"fmt"

	"github.com/jamascrorpJS/money-counter/internal/domain"
	"gorm.io/gorm"
)

type UsersDB struct {
	db *gorm.DB
}

func UsersRepo(db *gorm.DB) *UsersDB {

	return &UsersDB{
		db: db,
	}
}

func (c *UsersDB) Create(name string, surname string, email string, password string, image string) (int, error) {
	user := &domain.Users{
		Name:     name,
		Surname:  surname,
		Email:    email,
		Password: password,
		Image:    image,
	}
	err := c.db.Create(&user).Error
	if err != nil {
		fmt.Print(err)
		return 0, err
	}
	return user.ID, nil
}

func (c *UsersDB) GetByID(id string) *domain.Users {
	users := &domain.Users{}
	err := c.db.First(users, id).Error
	if err != nil {
		fmt.Print(err)
	}
	return users
}
func (c *UsersDB) GetAll() *[]domain.Users {
	users := &[]domain.Users{}
	err := c.db.Find(users).Error
	if err != nil {
		fmt.Print(err)
	}
	return users
}
func (c *UsersDB) GetSelectsFields(s []string) *domain.Users {
	users := &domain.Users{}
	err := c.db.Select(s).Find(users).Error
	if err != nil {
		fmt.Print(err)
	}
	return users
}

func (c *UsersDB) UpdateFields(id string, field *domain.Users) *domain.Users {
	users := &domain.Users{}
	err := c.db.Model(users).Where("id = ?", id).Updates(field).Error
	if err != nil {
		fmt.Print(err)
	}
	return users
}

func (c *UsersDB) Delete(id string) *domain.Users {
	users := &domain.Users{}
	err := c.db.Delete(users, id).Error
	if err != nil {
		fmt.Print(err)
	}
	return users
}

func (c *UsersDB) ExistedUsers(email string) (*domain.Users, error) {
	users := &domain.Users{}
	err := c.db.Model(users).Where("email = ?", email).First(users).Error
	if err != nil {
		print(users, err)
		return nil, err
	}
	print(users, err)
	return users, nil
}
