package helper

import (
	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	if data == nil {
		c.JSON(statusCode, gin.H{
			"status":  getStatusString(statusCode),
			"message": message,
		})
	}else{
		c.JSON(statusCode, gin.H{
			"status":  getStatusString(statusCode),
			"message": message,
			"data":    data,
		})
	}
}

func getStatusString(statusCode int) string {
	if statusCode >= 200 && statusCode < 300 {
		return "success"
	}
	return "error"
}
