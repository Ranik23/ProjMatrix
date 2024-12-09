package main

import (
	"ProjMatrix/internal/api/routes"
	"ProjMatrix/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
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

	router.Static("/assets", "../frontend/assets")
	router.Static("/css", "../frontend/css")
	router.Static("/js", "../frontend/js")
	router.LoadHTMLGlob("../frontend/views/*")

	routes.RegisterHTMLRoutes(router)
	routes.RegisterAPIRoutes(router)

	addr := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("Starting the %s server on %d", cfg.App.Name, cfg.App.Port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
