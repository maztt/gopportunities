package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
