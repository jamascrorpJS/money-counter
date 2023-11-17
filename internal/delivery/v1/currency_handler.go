package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/internal/repository"
	"github.com/jamascrorpJS/money-counter/internal/service"
	"github.com/jamascrorpJS/money-counter/pkg/cache"
	"gorm.io/gorm"
)

type CurrencyHandler struct {
	service *service.CurrencyService
}

func GetServiceCurrency(db *gorm.DB, cache *cache.Redis) *CurrencyHandler {
	return &CurrencyHandler{
		service: service.Currency(repository.CreateRepository(db, cache)),
	}
}

func (cu *CurrencyHandler) GetAllCurrency(c *gin.Context) {
	c.JSON(200, cu.service.GetAll())
}

func (cu *CurrencyHandler) GetByIDCurrency(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, cu.service.GetByID(id))
}

func (cu *CurrencyHandler) UpdateFieldsCurrency(c *gin.Context) {
	x := &domain.Currency{}
	c.ShouldBindJSON(&x)
	id := c.Param("id")
	c.JSON(200, cu.service.UpdateFields(id, x))
}

func (cu *CurrencyHandler) CreateCurrency(c *gin.Context) {
	x := domain.Currency{}
	c.ShouldBindJSON(&x)
	cu.service.Create(x.Name)
	c.JSON(200, x)
}
func (cu *CurrencyHandler) DeleteCurrency(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, cu.service.Delete(id))
}
