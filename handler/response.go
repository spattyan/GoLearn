package handler

import (
	"GoLearn/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, code int, message string) {
	context.Header("Content-type", "application/json")
	context.JSON(code, gin.H{
		"message": message,
		"errorCode": code,
	})
}

func sendSuccess(context *gin.Context, operation string, data interface{}) {
	context.Header("Content-type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %s successfull", operation),
		"data": data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	ErrorCode int `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string `json:"message"`
	Data schemas.OpeningResponse `json:"data"`
}