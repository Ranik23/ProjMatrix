// cmd/worker/main.go
package main

import (
	"ProjMatrix/internal/worker"
	"ProjMatrix/internal/worker/mw"
	"flag"
	"fmt"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "ProjMatrix/pkg/proto"
	"google.golang.org/grpc"
)

var (
	port     = flag.Int("port", 50051, "Worker port to listen on")
	workerId = flag.String("id", "", "Worker ID")
)

func main() {
	flag.Parse()

	if *workerId == "" {
		*workerId = fmt.Sprintf("worker-%d", *port)
	}

	serverOpts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1000 * 1024 * 1024), // 20MB
		grpc.MaxSendMsgSize(1000 * 1024 * 1024), // 20MB
		grpc.UnaryInterceptor(mw.Logging),
	}

	grpcServer := grpc.NewServer(serverOpts...)

	// Create worker processor (implement your specific job processing logic)

	// Create worker service
	workerService := worker.NewWorkerService(*workerId)

	// Register worker service
	reflection.Register(grpcServer)
	pb.RegisterWorkerServiceServer(grpcServer, workerService)

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
