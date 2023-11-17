package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/internal/repository"
	"github.com/jamascrorpJS/money-counter/internal/service"
	"github.com/jamascrorpJS/money-counter/pkg/cache"
	"gorm.io/gorm"
)

type Handler struct {
	service *service.CategoryService
}

func GetService(db *gorm.DB, cache *cache.Redis) *Handler {
	return &Handler{service: service.Category(repository.CreateRepository(db, cache))}
}

func (h *Handler) GetAll(c *gin.Context) {
	c.JSON(200, h.service.GetAll())
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, h.service.GetByID(id))
}

func (h *Handler) UpdateFields(c *gin.Context) {
	x := &domain.Category{}
	c.ShouldBindJSON(&x)
	id := c.Param("id")
	c.JSON(200, h.service.UpdateFields(id, x))
}

func (h *Handler) Create(c *gin.Context) {
	x := domain.CategoryDTO{}
	c.ShouldBindJSON(&x)
	h.service.Create(x.Name)
	c.JSON(200, x)
}
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, h.service.Delete(id))
}
