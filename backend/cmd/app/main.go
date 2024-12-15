package main

import (
	"ProjMatrix/internal/api/routes"
	"ProjMatrix/internal/config"
	"ProjMatrix/internal/entity"
	"ProjMatrix/pkg/proto"
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const worker_host = "localhost:7000"
const worker2_host = "localhost:7001"

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
	conn, err := grpc.NewClient(worker_host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn2, err := grpc.NewClient(worker2_host, grpc.WithTransportCredentials(insecure.NewCredentials()))

	workerClient := proto.NewWorkerServiceClient(conn)
	workerClient2 := proto.NewWorkerServiceClient(conn2)

	var WorkerClients entity.WorkersClient
	WorkerClients.FirstWorker.Client = workerClient
	WorkerClients.SecondWorker.Client = workerClient2

	routes.RegisterHTMLRoutes(router, WorkerClients)
	routes.RegisterAPIRoutes(router, WorkerClients)

	if err != nil {
		log.Fatalf("failed to create grpc client: %v", err)
	}
	defer conn.Close()

	addr := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("Starting the %s server on %d", cfg.App.Name, cfg.App.Port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
