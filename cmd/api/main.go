package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jozzhuve/technical-challenge-routing/internal/algorithm"
	"github.com/jozzhuve/technical-challenge-routing/internal/application"
	"github.com/jozzhuve/technical-challenge-routing/internal/httpapi"
)

// main configura las dependencias y levanta el servidor HTTP del servicio de rutas.
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	dijkstra := algorithm.NewDijkstra()
	service := application.NewCalculateOptimalRouteService(dijkstra)
	handler := httpapi.NewRouteHandler(service)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           httpapi.NewServer(handler),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	slog.Info("servicio de rutas iniciado", "port", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("el servidor terminó con error", "error", err)
		os.Exit(1)
	}
}
