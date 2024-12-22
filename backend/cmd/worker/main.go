// cmd/worker/main.go
package main

import (
	"ProjMatrix/internal/worker"
	"ProjMatrix/internal/worker/mw"
	"ProjMatrix/pkg/repository"
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "ProjMatrix/pkg/proto"
	"google.golang.org/grpc"
)

const psqlDSN = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

var (
	port     = flag.Int("port", 50051, "Worker port to listen on")
	workerId = flag.String("id", "", "Worker ID")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	if *workerId == "" {
		*workerId = fmt.Sprintf("worker-%d", *port)
	}

	serverOpts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1000 * 1024 * 1024), // 20MB
		grpc.MaxSendMsgSize(1000 * 1024 * 1024), // 20MB
		grpc.UnaryInterceptor(mw.Logging),
	}

	grpcServer := grpc.NewServer(serverOpts...)

	con, err := pgxpool.New(ctx, psqlDSN)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	pgRepository := repository.NewPgRepository(con)
	// Create worker processor (implement your specific job processing logic)

	// Create worker service
	workerService := worker.NewWorkerService(*workerId, pgRepository)
	// Register worker service
	reflection.Register(grpcServer)
	pb.RegisterWorkerServiceServer(grpcServer, workerService)

	workerService.Wp.Start()

	defer func() {
		workerService.Wp.Wait()
		workerService.Wp.Start()
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stop
		log.Printf("Shutting down worker %s...", *workerId)
		grpcServer.GracefulStop()
	}()

	// Start server
	log.Printf("Worker %s starting on port %d", *workerId, *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
