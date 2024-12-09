package routes

import (
	"ProjMatrix/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterHTMLRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/results", func(c *gin.Context) {
		log.Printf("\nРЕЗУЛЬТАТЫ\n\n %v", entity.ResultOfCalculations)
		c.HTML(http.StatusOK, "results.html", gin.H{
			"OperationType":    entity.ResultOfCalculations.OperationType,
			"ResultMatrix":     entity.ResultOfCalculations.ResultMatrix,
			"TimeCalc":         entity.ResultOfCalculations.TimeCalc,
			"TimeParallelCalc": entity.ResultOfCalculations.TimeParallelCalc,
		})
	})
}
