package repository

import (
	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/pkg/cache"
	"gorm.io/gorm"
)

type Category interface {
	Create(name string) int
	GetByID(id string) *domain.Category
	GetAll() *[]domain.Category
	GetSelectsFields(s []string) *domain.Category
	UpdateFields(id string, field *domain.Category) *domain.Category
	Delete(id string) *domain.Category
}

type Cost interface {
	Create(user_id uint, category_id uint, currency_id uint, summary int) int
	GetByID(id string) *domain.Costs
	GetAll() *[]domain.Costs
	GetSelectsFields(s []string) *domain.Costs
	UpdateFields(id string, field *domain.Costs) *domain.Costs
	Delete(id string) *domain.Costs
	Report() error
}

type Currency interface {
	Create(name string) int
	GetByID(id string) *domain.Currency
	GetAll() *[]domain.Currency
	GetSelectsFields(s []string) *domain.Currency
	UpdateFields(id string, field *domain.Currency) *domain.Currency
	Delete(id string) *domain.Currency
}

type Income interface {
	Create(users_id uint, currency_id uint, summary int) int
	GetByID(id string) *domain.Incomes
	GetAll() *[]domain.Incomes
	GetSelectsFields(s []string) *domain.Incomes
	UpdateFields(id string, field *domain.Incomes) *domain.Incomes
	Delete(id string) *domain.Incomes
}

type RefreshToken interface {
	Create(users_id uint, token string, expires_at string) int
	GetByID(id int) *domain.RefreshTokens
	GetAll() *[]domain.RefreshTokens
	GetSelectsFields(s ...string) *domain.RefreshTokens
	UpdateFields(field *domain.RefreshTokens) *domain.RefreshTokens
	Delete(id int) *domain.RefreshTokens
}
type Users interface {
	Create(name string, surname string, email string, password string, image string) (int, error)
	GetByID(id string) *domain.Users
	GetAll() *[]domain.Users
	GetSelectsFields(s []string) *domain.Users
	UpdateFields(id string, field *domain.Users) *domain.Users
	Delete(id string) *domain.Users
	ExistedUsers(email string) (*domain.Users, error)
}
type Repositories struct {
	Category     Category
	Cost         Cost
	Currency     Currency
	Income       Income
	RefreshToken RefreshToken
	Users        Users
}

func CreateRepository(db *gorm.DB, cache *cache.Redis) *Repositories {
	return &Repositories{
		Category:     CategoryRepo(db, cache),
		Cost:         CostRepo(db),
		Currency:     CurrencyRepo(db),
		Income:       IncomeRepo(db),
		RefreshToken: RefreshTokensRepo(db),
		Users:        UsersRepo(db),
	}
}
