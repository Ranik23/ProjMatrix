package main

import (
	"ProjMatrix/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

	// Тестовый маршрут
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})

	// Запуск сервера
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("Starting the %s server on %d", cfg.App.Name, cfg.App.Port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
