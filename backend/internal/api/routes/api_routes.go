package routes

import (
	"ProjMatrix/internal/api/handlers/request"
	"ProjMatrix/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterAPIRoutes(router *gin.Engine, workerClient *entity.WorkersClient) {
	router.POST("/api/submit", func(c *gin.Context) {
		var rawData map[string]interface{}

		if err := c.ShouldBindJSON(&rawData); err != nil {
			log.Println("JSON binding error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		if err := request.ProcessRequest(c, rawData, workerClient); err != nil {
			log.Println("Request processing error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
}
