package controllers


import (
	"examen/serverUno/application"
	"examen/serverUno/domain"
	"examen/serverUno/infraestructure"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreateProductHandlerChunked - Long Polling para creación de productos
func CreateProductHandler(c *gin.Context) {
	var product domain.Products

	// Decodificar el JSON del cuerpo de la solicitud
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	// Inicializar la BD y el caso de uso
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewCreateProduct(repo)

	// Configurar encabezados para chunked
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	encoder := json.NewEncoder(c.Writer)

	// Enviar primer chunk indicando que la creación ha comenzado
	encoder.Encode(gin.H{"status": "Iniciando la creación del producto"})
	c.Writer.Flush()
	time.Sleep(1 * time.Second)  // Simulamos latencia

	// Ejecutar la lógica de creación del producto
	if err := useCase.Execute(product); err != nil {
		encoder.Encode(gin.H{"error": "Error al guardar el producto"})
		c.Writer.Flush()
		return
	}

	// Enviar segundo chunk con éxito
	encoder.Encode(gin.H{"message": "Producto creado con éxito"})
	c.Writer.Flush()
}
