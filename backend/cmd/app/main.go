package main

import (
	"ProjMatrix/internal/api/routes"
	"ProjMatrix/internal/config"
	"ProjMatrix/internal/entity"
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	gob.Register(entity.CalculationResult{})
}

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

	// Хранилище для сессий (cookie-based)
	store := cookie.NewStore([]byte("champ_and_anton_key"))
	store.Options(sessions.Options{
		MaxAge:   3600,  // время жизни сессии
		Path:     "/",   // доступ ко всей проге
		HttpOnly: true,  // cookie недоступны из JS
		Secure:   false, // true ставить для https в production
	})

	// добавили middleware для работы с сессиями
	router.Use(sessions.Sessions("champ_and_anton_key", store))

	routes.RegisterHTMLRoutes(router)
	routes.RegisterAPIRoutes(router)

	addr := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("Starting the %s server on %d", cfg.App.Name, cfg.App.Port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
