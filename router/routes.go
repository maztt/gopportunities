package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maztt/gopportunities/handler"
)

// "*" is a pointer; != "pass by value"; works like you are giving an address for the code to access
func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// gin.H works as a wrapper for the object
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
