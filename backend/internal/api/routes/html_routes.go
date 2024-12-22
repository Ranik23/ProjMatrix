package routes

import (
	"ProjMatrix/internal/converter"
	"ProjMatrix/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHTMLRoutes(router *gin.Engine, workerClient *entity.WorkersClient) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/results", func(c *gin.Context) {
		value, dur, err := workerClient.PgRepository.Get(c, workerClient.Session)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		result := converter.ByteToCalculationResult(value)
		c.HTML(http.StatusOK, "results.html", gin.H{
			"OperationType":    result.OperationType,
			"ResultMatrix":     result.ResultMatrix,
			"TimeCalc":         dur,
			"TimeParallelCalc": result.TimeParallelCalc,
		})

	})
}
