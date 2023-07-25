package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maztt/gopportunities/schemas"
)

func reportError(context *gin.Context, code int, msg string) {
	context.Header("Content-Type", "application/json")
	context.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func reportSuccess(context *gin.Context, op string, data interface{}) {
	context.Header("Content-Type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s succesfully completed", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}