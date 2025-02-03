package controllers

import (
	"encoding/json"
	"examen/serverUno/application"
	"examen/serverUno/infraestructure"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)




func GetProductHandler(c *gin.Context) {
	// Configurar la cabecera para Transfer-Encoding: chunked
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")

	
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewGetProduct(repo)

	products, err := useCase.Execute()
	if err != nil {
		http.Error(c.Writer, "No se pudo obtener los productos", http.StatusInternalServerError)
		return
	}

	// Crear un encoder para enviar datos en chunks
	encoder := json.NewEncoder(c.Writer)

	// Escribir los productos en chunks
	for _, product := range products {
		// Convertir el producto a JSON y enviarlo
		if err := encoder.Encode(product); err != nil {
			break
		}

		// Forzar la escritura del chunk y hacer una pausa
		c.Writer.Flush()
		time.Sleep(1 * time.Second) // Simular retardo de streaming
	}
}
