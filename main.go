package main

import (
	"fmt"
	"kubajaru/rest-api-example/config"
	"kubajaru/rest-api-example/controller"
	"kubajaru/rest-api-example/repository"
	"kubajaru/rest-api-example/service"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // loads from .env
	cfg := config.LoadConfig()

	// Create a JSON handler with the level variable
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: &cfg.LogLevel,
	})

	// Set default logger
	logger := slog.New(handler)
	slog.SetDefault(logger)

	repo := repository.NewTaskRepository()
	svc := service.NewTaskService(repo)
	ctrl := controller.NewTaskController(svc)
	ctrl.RegisterRoutes()

	addr := fmt.Sprintf(":%s", cfg.Port)
	slog.Info("Starting server", "port", cfg.Port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
