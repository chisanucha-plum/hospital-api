package middleware

import (
	"hospital-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const prefix = "Bearer "
		tokenStr := c.GetHeader("Authorization")

		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		if len(tokenStr) <= len(prefix) || tokenStr[:len(prefix)] != prefix {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenStr = tokenStr[len(prefix):]

		token, err := services.ValidateJWT(tokenStr)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("staff_id", int(claims["staff_id"].(float64)))
		c.Set("hospital_id", int(claims["hospital_id"].(float64)))
		c.Next()
	}
}
