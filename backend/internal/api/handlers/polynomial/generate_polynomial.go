package polynomial

import (
	"ProjMatrix/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
)

func handleGeneratedPolynomial(c *gin.Context, p *entity.Polynomial) {
	log.Printf("Обработка генерации полинома: %+v\n", *p)
	// Logic for generated input
}
