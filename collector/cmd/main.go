package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"

	"golang-course/api/proto"
	"golang-course/collector/internal/adapter"
	"golang-course/collector/internal/config"
	"golang-course/collector/internal/delivery"
	"golang-course/collector/internal/usecase"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	githubAdapter := adapter.NewGitHubAdapter()
	collectorUC := usecase.NewCollectorUseCase(githubAdapter)
	grpcController := delivery.NewGrpcController(collectorUC)

	port := ":" + cfg.GRPCPort

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterGithubServiceServer(grpcServer, grpcController)

	log.Printf("Collector service is running on port %s...", port)

	// Graceful shutdown

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("Collector: shutting down gracefully...")
		done := make(chan struct{})

		go func() {
			grpcServer.GracefulStop()
			close(done)
		}()

		select {
		case <-done:
			log.Println("Collector stopped successfully.")
		case <-time.After(time.Duration(cfg.ShutdownTimeoutSeconds) * time.Second):
			log.Println("Collector: Shutdown timeout reached. Forcing stop...")
			grpcServer.Stop()
		}
	case err := <-errCh:
		log.Fatalf("gRPC server failed: %v", err)
	}
}
