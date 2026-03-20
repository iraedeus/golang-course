package main

import (
	"log"
	"net"

	"golang-course/api/proto"
	"golang-course/internal/collector/adapter"
	"golang-course/internal/collector/handler"
	"golang-course/internal/collector/usecase"

	"google.golang.org/grpc"
)

func main() {
	githubAdapter := adapter.NewGitHubAdapter()
	collectorUC := usecase.NewCollectorUseCase(githubAdapter)
	grpcHandler := handler.NewGrpcHandler(collectorUC)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterGithubServiceServer(grpcServer, grpcHandler)

	log.Println("Collector service is running on port :50051...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
