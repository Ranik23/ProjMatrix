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
		err := handleManualPolynomial(c, p)
		if err != nil {
			log.Printf("Ошибка при обработке вычислений: %w\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	case "generate-polynomial":
		err := handleGeneratedPolynomial(c, p)
		if err != nil {
			log.Printf("Ошибка при обработке вычислений: %w\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	default:
		log.Printf("Неизвестный operationType для полинома: %s\n", operationType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown operationType"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": p})
}
