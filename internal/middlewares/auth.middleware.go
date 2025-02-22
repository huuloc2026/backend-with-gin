package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/huuloc2026/go-backend/pkg/response"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.ErrorResponse(c, response.ErrorCodeInvalidToken, nil)
			c.Abort()
			return
		}

		// TODO: Xác thực token JWT ở đây

		c.Next()
	}
}
