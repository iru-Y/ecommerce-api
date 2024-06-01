package shared

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSucess(context *gin.Context, operation string, data interface{}) {
	context.Header("Content-type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from controller: %s successfully", operation),
		"data":    data,
	})
}

func ResponseError(context *gin.Context, code int, message string) {
	context.Header("Content-type", "application/json")
	context.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}
