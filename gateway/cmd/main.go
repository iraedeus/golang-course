package main

import (
	"log"
	"net/http"

	docs "golang-course/api/swagger"
	"golang-course/gateway/internal/adapter"
	"golang-course/gateway/internal/config"
	"golang-course/gateway/internal/delivery"
	"golang-course/gateway/internal/usecase"

	httpSwagger "github.com/swaggo/http-swagger"
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

	http.HandleFunc("/repo", httpHandler.GetRepository)
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Printf("Gateway service is running on port %s...", port)
	log.Printf("Collector address: %s", collectorAddr)
	log.Printf("Swagger documentation is available at http://localhost:%s/swagger/index.html", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
