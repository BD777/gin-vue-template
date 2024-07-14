package middleware

import (
	"gin-vue-template/pkg/utils/ctxmeta"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (m *Middleware) InitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logid := uuid.New().String()
		c.Set(string(ctxmeta.LogIDKey), logid)

		c.Next()
	}
}
