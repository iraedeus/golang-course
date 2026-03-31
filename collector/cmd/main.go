package main

import (
	"log"
	"net"

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

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
