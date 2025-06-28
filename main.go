package main

import (
	"kubajaru/rest-api-example/controller"
	"kubajaru/rest-api-example/repository"
	"kubajaru/rest-api-example/service"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// Define a log level variable (can be updated at runtime)
	var levelVar slog.LevelVar
	levelVar.Set(slog.LevelInfo) // Default level (can be set to LevelDebug, etc.)

	// Create a JSON handler with the level variable
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: &levelVar,
	})

	// Set default logger
	logger := slog.New(handler)
	slog.SetDefault(logger)

	repo := repository.NewTaskRepository()
	svc := service.NewTaskService(repo)
	ctrl := controller.NewTaskController(svc)
	ctrl.RegisterRoutes()

	slog.Info("Starting server", "port", 8080)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
