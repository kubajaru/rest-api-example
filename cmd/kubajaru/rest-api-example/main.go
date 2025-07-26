package main

import (
	"fmt"
	"kubajaru/rest-api-example/internal/config"
	"kubajaru/rest-api-example/internal/controller"
	"kubajaru/rest-api-example/internal/model"
	"kubajaru/rest-api-example/internal/repository"
	"kubajaru/rest-api-example/internal/service"
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

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	slog.Info("Starting server", "port", cfg.Port)
	err := http.ListenAndServe(addr, setupRouter())
	if err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}

func setupRouter() *http.ServeMux {
	// Initalize router
	mux := http.NewServeMux()

	// Initalize Tasks controller
	taskRepo := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskController := controller.NewTaskController(taskService)
	taskController.RegisterRoutes(mux)

	// This should not be here, testing of test only
	task := model.Task{
		Title: "sample",
		Done:  false,
	}

	taskService.Create(task)

	return mux
}
