package linear

import (
	"ProjMatrix/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleLinearForm(c *gin.Context, l *entity.LinearForm, operationType string) {
	switch operationType {
	case "manual-linear-form":
		err := handleManualLinearForm(c, l)
		if err != nil {
			log.Printf("Ошибка при обработке вычислений: %w\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	case "generate-linear-form":
		err := handleGeneratedLinearForm(c, l)
		if err != nil {
			log.Printf("Ошибка при обработке вычислений: %w\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	default:
		log.Printf("Неизвестный operationType для линейной формы: %s\n", operationType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown operationType"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": *l})
}
