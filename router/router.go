package router

import "github.com/gin-gonic/gin"

// "Any entity" capitalized is automatically exportable
func Initialize() {
	// Initializes the Router variable:
	router := gin.Default() // "r"outer is an ephemere variable (convention)

	// Initialize Routes
	initializeRoutes(router)

	// Running the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
