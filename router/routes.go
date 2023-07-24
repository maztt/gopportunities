package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/maztt/gopportunities/docs"
	"github.com/maztt/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// "*" is a pointer; != "pass by value"; works like you are giving an address for the code to access
func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		// gin.H works as a wrapper for the object
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	// Initialize Swagger
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
