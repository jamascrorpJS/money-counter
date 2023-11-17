package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/pkg/cache"
	"gorm.io/gorm"
)

type CategoryDB struct {
	db    *gorm.DB
	cache *cache.Redis
}

func CategoryRepo(db *gorm.DB, cache *cache.Redis) *CategoryDB {

	return &CategoryDB{
		db:    db,
		cache: cache,
	}
}

func (c *CategoryDB) Create(name string) int {
	res := c.db.Create(&domain.Category{
		Name: name,
	})
	if res.Error != nil {
		fmt.Print(res.Error)
	}
	return int(res.RowsAffected)
}

func (c *CategoryDB) GetByID(id string) *domain.Category {
	category := &domain.Category{}
	err := c.db.First(category, id).Error
	if err != nil {
		fmt.Print(err)
	}
	return category
}

func (c *CategoryDB) GetSelectsFields(s []string) *domain.Category {
	category := &domain.Category{}
	err := c.db.Select(s).Find(category).Error
	if err != nil {
		fmt.Print(err)
	}
	return category
}

func (c *CategoryDB) UpdateFields(id string, field *domain.Category) *domain.Category {
	category := &domain.Category{}
	err := c.db.Model(category).Where("id = ?", id).Updates(field).Error
	if err != nil {
		fmt.Print(err)
		return category
	}
	jsonMarshal, ers := json.Marshal(category)
	if ers != nil {

	}
	errorx := c.cache.SetKey(context.Background(), "category", jsonMarshal, 1000*time.Second)
	if errorx != nil {

	}
	return category
}

func (c *CategoryDB) Delete(id string) *domain.Category {
	category := &domain.Category{}
	err := c.db.Delete(category, id).Error
	if err != nil {
		fmt.Print(err)
	}

	c.cache.Del(context.Background(), "category")
	return category
}

func (c *CategoryDB) GetAll() *[]domain.Category {
	category := &[]domain.Category{}
	red, err := c.cache.GetValue(context.Background(), "category")
	if err == nil {
		jsonerr := json.Unmarshal([]byte(red), &category)
		if jsonerr != nil {

		}
		fmt.Print("cache")
		return category
	}
	dbErr := c.db.Find(category).Error
	if dbErr != nil {

	}
	jsonMarshal, ers := json.Marshal(category)
	if ers != nil {

	}
	errorx := c.cache.SetKey(context.Background(), "category", jsonMarshal, 1000*time.Second)
	if errorx != nil {

	}
	fmt.Print("no cache")
	return category
}
