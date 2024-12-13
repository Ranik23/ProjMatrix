package routes

import (
	"ProjMatrix/internal/api/handlers/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterAPIRoutes(router *gin.Engine) {
	router.POST("/api/submit", func(c *gin.Context) {
		var rawData map[string]interface{}

		if err := c.ShouldBindJSON(&rawData); err != nil {
			log.Println("JSON binding error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		if err := request.ProcessRequest(c, rawData); err != nil {
			log.Println("Request processing error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
}
