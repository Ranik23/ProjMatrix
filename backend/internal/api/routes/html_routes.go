package routes

import (
	"ProjMatrix/internal/entity"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterHTMLRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/results", func(c *gin.Context) {
		session := sessions.Default(c)
		value := session.Get("calculationResult")

		if result, ok := value.(entity.CalculationResult); ok {
			c.HTML(http.StatusOK, "results.html", gin.H{
				"OperationType":    result.OperationType,
				"ResultMatrix":     result.ResultMatrix,
				"TimeCalc":         result.TimeCalc,
				"TimeParallelCalc": result.TimeParallelCalc,
			})
		} else {
			log.Printf("Error getting the result: %+v \n", value)
			c.JSON(http.StatusNotFound, gin.H{"error": "The result was not found. Perform the calculation again."})
		}
	})
}
