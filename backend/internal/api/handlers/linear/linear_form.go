package linear

import (
	"ProjMatrix/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleLinearForm(c *gin.Context, l *entity.LinearForm, operationType string) {
	// Differentiate manual input vs generation
	switch operationType {
	case "manual-linear-form":
		handleManualLinearForm(l)
	case "generate-linear-form":
		handleGeneratedLinearForm(l)
	default:
		log.Printf("Неизвестный operationType для линейной формы: %s\n", operationType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown operationType"})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": *l})
}
