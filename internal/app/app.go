package app

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jamascrorpJS/money-counter/internal/delivery/v1"
	"github.com/jamascrorpJS/money-counter/pkg/cache"
	"github.com/jamascrorpJS/money-counter/pkg/database"
)

func StartServer() {
	c := gin.Default()
	db := database.NewClient()
	cache := cache.NewRedisClient()
	categ := c.Group("/category")
	{
		handler := v1.GetService(db, cache)
		categ.GET("/", handler.GetAll)
		categ.GET("/:id", handler.GetByID)
		categ.POST("/", handler.Create)
		categ.PUT("/:id", handler.UpdateFields)
		categ.DELETE("/:id", handler.Delete)
	}
	costs := c.Group("/costs")
	{
		handler := v1.GetServiceCosts(db, cache)
		costs.GET("/", handler.GetAllCosts)
		costs.GET("/:id", handler.GetByIDCosts)
		costs.POST("/", handler.CreateCosts)
		costs.PUT("/:id", handler.UpdateFieldsCosts)
		costs.DELETE("/:id", handler.DeleteCosts)
		costs.GET("/report", handler.Report)
		costs.GET("/report/loads", handler.Download)
	}

	curency := c.Group("/currency")
	{
		handler := v1.GetServiceCurrency(db, cache)
		curency.GET("/", handler.GetAllCurrency)
		curency.GET("/:id", handler.GetByIDCurrency)
		curency.POST("/", handler.CreateCurrency)
		curency.PUT("/:id", handler.UpdateFieldsCurrency)
		curency.DELETE("/:id", handler.DeleteCurrency)

	}

	users := c.Group("/users")

	{
		handler := v1.GetServiceUsers(db, cache)
		users.GET("/", handler.GetAllUsers)
		users.GET("/:id", handler.GetByIDUsers)
		users.POST("/", handler.CreateUsers)
		users.PUT("/:id", handler.UpdateFieldsUsers)
		users.DELETE("/:id", handler.DeleteUsers)
		users.POST("/login", handler.LoginUser)
	}
	c.Run("localhost:8080")
}
