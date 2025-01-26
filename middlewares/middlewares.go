package middlewares

import (
	"net/http"
	"strings"

	"sidita-be/utils/helper"
	"sidita-be/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) > 0 && !strings.HasPrefix(authHeader, "Bearer ") {
			authHeader = "Bearer " + authHeader
			c.Request.Header.Set("Authorization", authHeader)
		}

		err := token.TokenValid(c)
		if err != nil {
			helper.SendResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}