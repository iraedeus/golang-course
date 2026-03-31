package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

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
	defer grpcClient.Close()

	gatewayUC := usecase.NewGatewayUseCase(grpcClient)
	httpHandler := delivery.NewHTTPController(gatewayUC)

	router := delivery.NewRouter(httpHandler)

	srv := &http.Server{Addr: ":" + port, Handler: router}

	log.Printf("Gateway service is running on port %s...", port)
	log.Printf("Swagger documentation is available at http://localhost:%s/swagger/index.html", port)

	// Graceful shutdown

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("Gateway: shutting down gracefully...")
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(cfg.ShutdownTimeoutSeconds)*time.Second,
		)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Println("Gateway: Shutdown timeout reached. Forcing stop...")
			srv.Close()
			return
		}

		log.Println("Gateway service stopped successfully")
	case err := <-errCh:
		log.Fatalf("HTTP server failed: %v", err)
	}
}
