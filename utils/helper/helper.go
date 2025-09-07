package helper

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
)

type PaginationResponse struct {
	//Message string                 `json:"message"`
	Data map[string]interface{} `json:"data"`
	Page Page                   `json:"page"`
	Message   string `json:"message"`
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code,omitempty"`
}

type Page struct {
	CurrentPage int   `json:"current_page" json:"current_page,omitempty"`
	TotalData   int64 `json:"total_data" json:"total_data,omitempty"`
	Limit       int   `json:"limit" json:"limit,omitempty"`
	TotalPage   int   `json:"total_page" json:"total_page,omitempty"`
}

func JSON(response interface{}, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("error", slog.Any("error", err))
		return
	}
}

func SuccessDataPaginate(msg, dataName string, data interface{}, page *Page, w http.ResponseWriter) {
	response := PaginationResponse{
		Message: msg,
		Code:    200,
		Data: map[string]interface{}{
			dataName: data,
		},
		Page: Page{
			Limit:       page.Limit,
			TotalPage:   page.TotalPage,
			TotalData:   page.TotalData,
			CurrentPage: page.CurrentPage,
		},
	}
	JSON(response, 200, w)
}

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
