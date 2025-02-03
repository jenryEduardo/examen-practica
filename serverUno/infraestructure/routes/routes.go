package routes

import (
	"examen/serverUno/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura las rutas con Gin
func SetupRoutes(router *gin.Engine) {
	// Rutas para productos
	router.POST("/products", controllers.CreateProductHandler)
	router.GET("/view-products", controllers.GetProductHandler)
	router.DELETE("/delete-products/:id", controllers.DeleteProductHandler)
	router.PUT("/update-products/:id", controllers.UpdateProductHandler)
}


