package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// "*" is a pointer; != "pass by value"; works like you are giving an address for the code to access
func initializeRoutes(router *gin.Engine) { 
	v1 := router.Group("/api/v1")
	{
		// gin.H works as a wrapper for the object
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "GET.opening",
			})
		})
		v1.POST("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "POST.opening",
			})
		})
		v1.DELETE("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "DELETE.opening",
			})
		})
		v1.PUT("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "PUT.opening",
			})
		})
		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "GET.openings",
			})
		})
	}
}
