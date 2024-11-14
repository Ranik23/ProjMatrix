package main

import (
	"ProjMatrix/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
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
