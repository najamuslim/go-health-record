package middleware

import (
	"fmt"
	"health-record/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	helper helpers.HelperInterface
}

type MiddlewareInterface interface {
	AuthMiddleware(c *gin.Context)
}

func NewMiddleware(helper helpers.HelperInterface) MiddlewareInterface {
	return &Middleware{helper}
}

func (m *Middleware) AuthMiddleware(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")

	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, "request token is missing or expired")
		c.Abort()
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, "invalid token format")
		c.Abort()
		return
	}

	// Validate JWT
	claims, err := m.helper.ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "invalid token")
		c.Abort()
		return
	}

	num, err := strconv.Atoi(claims.Subject)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.Set("user_id", num)
	c.Next()
}
