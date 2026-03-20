package main

import (
	"log"
	"net/http"

	"golang-course/internal/gateway/client"
	"golang-course/internal/gateway/handler"

	_ "golang-course/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           GitHub Info API
// @version         1.0
// @description     This is a sample server for fetching GitHub repo info.
// @host            localhost:8080
// @BasePath        /
func main() {
	collectorClient, err := client.NewCollectorClient("localhost:50051")
	if err != nil {
		log.Fatalf("could not connect to collector: %v", err)
	}

	h := handler.NewHttpHandler(collectorClient)
	http.HandleFunc("/repo", h.GetRepository)

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Println("Gateway service is running on port :8080...")
	log.Println("Swagger documentation is available at http://localhost:8080/swagger/index.html")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
