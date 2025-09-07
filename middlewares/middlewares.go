package middlewares

import (
	"net/http"
	"strings"

	"note-test/models"
	"note-test/utils/helper"
	"note-test/utils/token"

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

		userID, err := token.ExtractTokenID(c)

		if err != nil {
			helper.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
			return
		}

		err = models.CreateLog(&models.Log{
			EndPoint: c.FullPath(),
			Method:   c.Request.Method,
			UserID:   userID,
		})
		if err != nil {
			helper.SendResponse(c, http.StatusInternalServerError, "Failed to create log", nil)
			return
		}

		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
