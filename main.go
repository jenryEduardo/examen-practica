
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	productRoutes "examen/serverUno/infraestructure/routes"
	
)

func main() {
	
	router := gin.Default()

	// Configurar rutas de productos y usuarios
	productRoutes.SetupRoutes(router)
	

	port := ":8080"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))
}
