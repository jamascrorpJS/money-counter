package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jamascrorpJS/money-counter/internal/domain"
	"github.com/jamascrorpJS/money-counter/internal/repository"
	"github.com/jamascrorpJS/money-counter/internal/service"
	"github.com/jamascrorpJS/money-counter/pkg/cache"
	"gorm.io/gorm"

	jwt "github.com/jamascrorpJS/money-counter/pkg/token"
)

type UsersHandler struct {
	service *service.UsersService
}

func GetServiceUsers(db *gorm.DB, cache *cache.Redis) *UsersHandler {
	return &UsersHandler{
		service: service.Users(repository.CreateRepository(db, cache)),
	}
}

func (us *UsersHandler) GetAllUsers(c *gin.Context) {
	c.JSON(200, us.service.GetAll())
}
func (us *UsersHandler) GetByIDUsers(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, us.service.GetByID(id))
}

func (us *UsersHandler) UpdateFieldsUsers(c *gin.Context) {
	x := &domain.Users{}
	c.ShouldBindJSON(&x)
	id := c.Param("id")
	c.JSON(200, us.service.UpdateFields(id, x))
}
func (us *UsersHandler) CreateUsers(c *gin.Context) {
	x := domain.Users{}
	c.ShouldBindJSON(&x)
	user, err := us.service.Create(x.Name, x.Surname, x.Email, x.Password, x.Image)

	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
	jwt.Jw(user)
}
func (us *UsersHandler) LoginUser(c *gin.Context) {
	x := &domain.Users{}
	c.ShouldBindJSON(&x)
	jw, err := us.service.LoginUsers(x.Email, x.Password)
	print(err)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.SetCookie("access_token", jw, 10000, "/", "", false, true)
	c.JSON(200, gin.H{"access_token": jw})
}

func (us *UsersHandler) DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, us.service.Delete(id))
}

func (us *UsersHandler) ExistedUsers(c *gin.Context) {
	x := domain.Users{}
	c.ShouldBindJSON(&x)
	user, err := us.service.ExistedUsers(x.Email)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}
