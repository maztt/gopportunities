package router

import "github.com/gin-gonic/gin"

// "Qualquer entidade" que começar com a letra inicial maiúscula, é exportável
func Initialize() {
	// Inicializa a variável:
	router := gin.Default() // "r"outer é uma variável efêmera (convenção)

	// Define uma rota:
	router.GET("/ping", func(c *gin.Context) { // "c"ontext
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
