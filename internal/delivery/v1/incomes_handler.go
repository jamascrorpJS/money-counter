package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/internal/repository"
	"github.com/jamascrorpJS/money-counter/internal/service"
	"github.com/jamascrorpJS/money-counter/pkg/cache"
	"gorm.io/gorm"
)

type IncomesHandler struct {
	service *service.IncomesService
}

func getServiceIncomes(db *gorm.DB, cache *cache.Redis) *IncomesHandler {
	return &IncomesHandler{
		service: service.Incomes(repository.CreateRepository(db, cache)),
	}
}

func (ic *IncomesHandler) GetAllIncomes(c *gin.Context) {
	c.JSON(200, ic.service.GetAll())
}

func (ic *IncomesHandler) GetByIDIncomes(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, ic.service.GetByID(id))
}

func (ic *IncomesHandler) UpdateFieldsIncomes(c *gin.Context) {
	x := &domain.Incomes{}
	c.ShouldBindJSON(&x)
	id := c.Param("id")
	c.JSON(200, ic.service.UpdateFields(id, x))
}

func (ic *IncomesHandler) CreateIncomes(c *gin.Context) {
	x := domain.Incomes{}
	c.ShouldBindJSON(&x)
	ic.service.Create(x.UsersID, x.CurrencyID, x.Summary)
	c.JSON(200, x)
}
func (ic *IncomesHandler) DeleteIncomes(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, ic.service.Delete(id))
}
