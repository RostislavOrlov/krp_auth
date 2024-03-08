package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("старт работы middleware")
		authHeader := c.GetHeader("Authorization")
		logrus.Info("authorization header: ", authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization header",
			})
			c.Abort()
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization header format",
			})
			c.Abort()
			return
		}

		token := authHeaderParts[1]
		fmt.Println("Token:", token)
		c.Set("access_token", token)
		c.Next()
	}
}
