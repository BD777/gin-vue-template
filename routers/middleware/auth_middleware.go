package middleware

import (
	"gin-vue-template/pkg/constants"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		token, err := c.Cookie(constants.TokenKey)
		if err != nil {
			log.Printf("[AuthMiddleware] get cookie token error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		loginInfo, err := m.logic.GetLoginInfo(ctx, token)
		if err != nil {
			log.Printf("[AuthMiddleware] logic.GetLoginInfo error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set user info
		c.Set(constants.AuthKey, loginInfo)

		c.Next()
	}
}
