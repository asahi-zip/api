package middlewares

import (
	"net/http"
	"strings"

	"github.com/asahi-zip/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if len(token) == 0 || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		tokenValue := strings.TrimPrefix(token, "Bearer ")

		var user models.User
		if err := models.DB.Where("token = ?", tokenValue).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query database"})
			c.Abort()
			return
		}

		c.Next()
	}
}
