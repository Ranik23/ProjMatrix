package main

import (
	"ProjMatrix/internal/api/routes"
	"ProjMatrix/internal/config"
	"ProjMatrix/internal/entity"
	"ProjMatrix/pkg/proto"
	"ProjMatrix/pkg/repository"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const pg = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
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
	ctx := context.Background()
	// Режим Gin (всего их три, но будем использовать только два)
	if cfg.App.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.Default()

	con, err := pgxpool.New(ctx, pg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	pgRepository := repository.NewPgRepository(con)

	router.Static("/assets", "../frontend/assets")
	router.Static("/css", "../frontend/css")
	router.Static("/js", "../frontend/js")
	router.LoadHTMLGlob("../frontend/views/*")

	conn, err := grpc.NewClient(worker_host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn2, err := grpc.NewClient(worker2_host, grpc.WithTransportCredentials(insecure.NewCredentials()))

	workerClient := proto.NewWorkerServiceClient(conn)
	workerClient2 := proto.NewWorkerServiceClient(conn2)

	WorkerClients := entity.NewWorkersClient(workerClient, workerClient2, pgRepository, "")

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
