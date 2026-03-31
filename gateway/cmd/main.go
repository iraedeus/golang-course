package main

import (
	"log"
	"net/http"

	docs "golang-course/api/swagger"
	"golang-course/gateway/internal/adapter"
	"golang-course/gateway/internal/config"
	"golang-course/gateway/internal/delivery"
	"golang-course/gateway/internal/usecase"
)

// @title           GitHub Info API
// @version         1.0
// @description     This is a sample server for fetching GitHub repo info.
// @host            localhost:8080
// @BasePath        /
func main() {
	cfg := config.Load()
	port := cfg.HTTPPort
	collectorAddr := cfg.CollectorAddr

	docs.SwaggerInfo.Host = "localhost:" + port

	grpcClient, err := adapter.NewCollectorGrpcClient(collectorAddr)
	if err != nil {
		log.Fatalf("could not connect to collector: %v", err)
	}

	gatewayUC := usecase.NewGatewayUseCase(grpcClient)
	httpHandler := delivery.NewHttpHandler(gatewayUC)

	router := delivery.NewRouter(httpHandler)

	log.Printf("Gateway service is running on port %s...", port)
	log.Printf("Collector address: %s", collectorAddr)
	log.Printf("Swagger documentation is available at http://localhost:%s/swagger/index.html", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
