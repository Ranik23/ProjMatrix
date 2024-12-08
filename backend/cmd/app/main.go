package main

import (
	"ProjMatrix/internal/config"
	"ProjMatrix/internal/models"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig("/home/champ001/GolandProjects/ProjMatrix/backend/configs/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Режим Gin (всего их три, но будем использовать только два)
	if cfg.App.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.Default()

	// Настройка статичстичеких файлов
	router.Static("/assets", "/home/champ001/GolandProjects/ProjMatrix/frontend/assets")
	router.Static("/css", "/home/champ001/GolandProjects/ProjMatrix/frontend/css")
	router.Static("/js", "/home/champ001/GolandProjects/ProjMatrix/frontend/js")

	// Настройка марштрутов для рендернига HTML-страниц
	router.LoadHTMLGlob("/home/champ001/GolandProjects/ProjMatrix/frontend/views/*")

	// Маршрут для главной страницы
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//router.POST("/api/generate-linear-form", func(c *gin.Context) {
	//	var requestData struct {
	//		GenerationType string `json:"generationType"`
	//		MatrixCount    int    `json:"matrixCount"`
	//		MatrixSize     struct {
	//			Rows    int `json:"rows"`
	//			Columns int `json:"columns"`
	//		} `json:"matrixSize"`
	//	}
	//
	//	// Попытка привязать JSON к структуре
	//	if err := c.ShouldBindJSON(&requestData); err != nil {
	//		log.Println("Ошибка привязки JSON:", err)
	//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
	//		return
	//	}
	//
	//	// Логируем полученные данные
	//	log.Printf("\nПолучены данные для генерации линейной формы:\n %+v\n\n", requestData)
	//
	//	// Отправляем успешный ответ
	//	c.JSON(http.StatusOK, gin.H{"status": "success", "data": requestData})
	//})
	//
	//router.POST("/api/generate-polynomial", func(c *gin.Context) {
	//	var requestData struct {
	//		GenerationType string `json:"generationType"`
	//		MatrixSize     struct {
	//			Rows    int `json:"rows"`
	//			Columns int `json:"columns"`
	//		} `json:"matrixSize"`
	//		Degree int `json:"degree"`
	//	}
	//
	//	if err := c.ShouldBindJSON(&requestData); err != nil {
	//		log.Println("Ошибка привязки JSON:", err)
	//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	//		return
	//	}
	//
	//	// Логируем полученные данные
	//	log.Printf("\nПолучены данные для генерации полинома:\n %+v\n\n", requestData)
	//
	//	// Ответ клиенту
	//	c.JSON(http.StatusOK, gin.H{"status": "success", "data": requestData})
	//})
	//
	//router.POST("/api/manual-polynomial", func(c *gin.Context) {
	//	log.Println("POST запрос получен на /api/manual-polynomial")
	//
	//	var requestData struct {
	//		MatrixSize struct {
	//			Rows    int `json:"rows"`
	//			Columns int `json:"columns"`
	//		} `json:"matrixSize"`
	//		Matrix       []float64 `json:"matrix"`
	//		Degree       int       `json:"degree"`
	//		Coefficients []float64 `json:"coefficients"`
	//	}
	//
	//	if err := c.ShouldBindJSON(&requestData); err != nil {
	//		log.Println("Ошибка привязки JSON:", err)
	//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	//		return
	//	}
	//
	//	log.Printf("Получены данные: %+v\n", requestData)
	//	c.JSON(http.StatusOK, gin.H{"status": "success", "data": requestData})
	//})
	//
	//router.POST("/api/manual-linear-form", func(c *gin.Context) {
	//	var requestData struct {
	//		MatrixCount int `json:"matrixCount"`
	//		MatrixSize  struct {
	//			Rows    int `json:"rows"`
	//			Columns int `json:"columns"`
	//		} `json:"matrixSize"`
	//		Matrices     [][]float64 `json:"matrices"`
	//		Coefficients []float64   `json:"coefficients"`
	//	}
	//
	//	if err := c.ShouldBindJSON(&requestData); err != nil {
	//		log.Println("Ошибка привязки JSON:", err)
	//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	//		return
	//	}
	//
	//	// Логируем полученные данные
	//	log.Printf("\nПолучены данные для ручного заполнения линейной формы:\n %+v\n\n", requestData)
	//
	//	// Ответ клиенту
	//	c.JSON(http.StatusOK, gin.H{"status": "success", "data": requestData})
	//})

	router.POST("/api/submit", func(c *gin.Context) {
		var baseRequest struct {
			OperationType string `json:"operationType"`
		}

		// Привязываем JSON к структуре baseRequest
		if err := c.ShouldBindJSON(&baseRequest); err != nil {
			log.Println("Ошибка привязки JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		// Логируем тип операции
		log.Printf("Получен запрос с operationType: %s\n", baseRequest.OperationType)

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Обрабатываем запрос в зависимости от operationType
		switch baseRequest.OperationType {
		case "manual-polynomial", "generate-polynomial":
			handlePolynomial(c)
		case "manual-linear-form", "generate-linear-form":
			handleLinearForm(c)
		default:
			log.Printf("Неизвестный operationType: %s\n", baseRequest.OperationType)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown operationType"})
		}
	})

	// Маршрут для страницы результатов
	router.GET("/results", func(c *gin.Context) {
		c.HTML(http.StatusOK, "results.html", nil)
	})

	// Запуск сервера
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("Starting the %s server on %d", cfg.App.Name, cfg.App.Port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func handlePolynomial(c *gin.Context) {
	var polynomialRequest models.Polynomial

	if err := c.ShouldBindJSON(&polynomialRequest); err != nil {
		log.Println("Ошибка привязки JSON к Polynomial:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Polynomial JSON"})
		return
	}

	switch polynomialRequest.OperationType {
	case "manual-polynomial":
		handleManualPolynomial(polynomialRequest)
	case "generate-polynomial":
		handleGeneratedPolynomial(polynomialRequest)
	default:
		log.Printf("Неизвестный operationType для полинома: %s\n", polynomialRequest.OperationType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown operationType"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": polynomialRequest})
}

func handleManualPolynomial(request models.Polynomial) {
	log.Printf("Обработка ручного ввода полинома: %+v\n", request)
	// Logic for manual input
}

func handleGeneratedPolynomial(request models.Polynomial) {
	log.Printf("Обработка генерации полинома: %+v\n", request)
	// Logic for generated input
}

func handleLinearForm(c *gin.Context) {
	var linearFormRequest models.LinearForm

	// Bind JSON to LinearForm structure
	if err := c.ShouldBindJSON(&linearFormRequest); err != nil {
		log.Println("Ошибка привязки JSON к LinearForm:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Linear Form JSON"})
		return
	}

	// Differentiate manual input vs generation
	switch linearFormRequest.OperationType {
	case "manual-linear-form":
		handleManualLinearForm(linearFormRequest)
	case "generate-linear-form":
		handleGeneratedLinearForm(linearFormRequest)
	default:
		log.Printf("Неизвестный operationType для линейной формы: %s\n", linearFormRequest.OperationType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown operationType"})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": linearFormRequest})
}

func handleManualLinearForm(request models.LinearForm) {
	log.Printf("Обработка ручного ввода линейной формы: %+v\n", request)
	// Logic for manual input
}

func handleGeneratedLinearForm(request models.LinearForm) {
	log.Printf("Обработка генерации линейной формы: %+v\n", request)
	// Logic for generated input
}
