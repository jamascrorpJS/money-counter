package v1

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/internal/repository"
	"github.com/jamascrorpJS/money-counter/internal/service"
	"github.com/jamascrorpJS/money-counter/pkg/cache"
	"gorm.io/gorm"
)

type CostHandler struct {
	service *service.CostsService
}

func GetServiceCosts(db *gorm.DB, cache *cache.Redis) *CostHandler {
	return &CostHandler{
		service: service.Costs(repository.CreateRepository(db, cache)),
	}
}

func (co *CostHandler) GetAllCosts(c *gin.Context) {
	c.JSON(200, co.service.GetAll())
}

func (co *CostHandler) GetByIDCosts(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, co.service.GetByID(id))
}

func (co *CostHandler) UpdateFieldsCosts(c *gin.Context) {
	x := &domain.Costs{}
	c.ShouldBindJSON(&x)
	id := c.Param("id")
	c.JSON(200, co.service.UpdateFields(id, x))
}

func (co *CostHandler) CreateCosts(c *gin.Context) {
	x := domain.Costs{}
	c.ShouldBindJSON(&x)
	co.service.Create(x.UsersID, x.CategoryID, x.CurrencyID, x.Summary)
	c.JSON(200, x)
}
func (co *CostHandler) DeleteCosts(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, co.service.Delete(id))
}

func (co *CostHandler) Report(c *gin.Context) {
	e := co.service.Report()
	if e != nil {
		c.JSON(400, gin.H{"error": ""})
	}
	c.JSON(200, gin.H{})
}

func (co *CostHandler) Download(c *gin.Context) {
	filePath := "Book1.xlsx"

	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка открытия файла: %s", err))
		return
	}
	defer file.Close()

	// Устанавливаем заголовок Content-Disposition для указания имени файла
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "Book1.xlsx"))

	// Копируем содержимое файла в ответ
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка отправки файла: %s", err))
		return
	}
}
