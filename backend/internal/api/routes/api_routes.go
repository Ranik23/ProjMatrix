package routes

import (
	handLin "ProjMatrix/internal/api/handlers/linear"
	handPol "ProjMatrix/internal/api/handlers/polynomial"
	"ProjMatrix/internal/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterAPIRoutes(router *gin.Engine) {
	router.POST("/api/submit", func(c *gin.Context) {
		var rawData map[string]interface{}

		// Привязываем JSON к структуре baseRequest
		if err := c.ShouldBindJSON(&rawData); err != nil {
			log.Println("Ошибка привязки JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		operationType, ok := rawData["operationType"].(string)
		if !ok {
			log.Println("Поле operationType отсутствует или имеет некорректный тип")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing operationType"})
			return
		}

		switch operationType {
		case "manual-polynomial", "generate-polynomial":
			var polynomial entity.Polynomial
			if err := mapToStruct(rawData, &polynomial); err != nil {
				log.Println("Ошибка привязки к Polynomial:", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data for Polynomial"})
				return
			}
			handPol.HandlePolynomial(c, &polynomial, operationType)
			log.Println("Получена структура Polynomial:", polynomial)

		case "manual-linear-form", "generate-linear-form":
			var linearForm entity.LinearForm
			if err := mapToStruct(rawData, &linearForm); err != nil {
				log.Println("Ошибка привязки к LinearForm:", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data for LinearForm"})
				return
			}
			handLin.HandleLinearForm(c, &linearForm, operationType)
			log.Println("Получена структура LinearForm:", linearForm)

		default:
			log.Println("Неизвестный operationType:", operationType)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown operationType"})
			return
		}
	})
}

func mapToStruct(data map[string]interface{}, result interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, result)
}
