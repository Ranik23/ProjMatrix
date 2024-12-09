package polynomial

import (
	"ProjMatrix/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandlePolynomial(c *gin.Context, p *entity.Polynomial, operationType string) {
	switch operationType {
	case "manual-polynomial":
		handleManualPolynomial(c, p)
	case "generate-polynomial":
		handleGeneratedPolynomial(c, p)
	default:
		log.Printf("Неизвестный operationType для полинома: %s\n", operationType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown operationType"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": p})
}
