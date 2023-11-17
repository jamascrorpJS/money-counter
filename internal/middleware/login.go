package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt "github.com/jamascrorpJS/money-counter/pkg/token"
)

func CheckToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		token, _ := c.Cookie("access_token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Требуется токен для доступа"})
			c.Abort()
			return
		}
		id, err := jwt.Jws(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		fmt.Print(id)
		// user := "someuser"
		// c.Set("user", user)

		// c.Next()
	}
}
