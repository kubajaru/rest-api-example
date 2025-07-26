package controller

import (
	"encoding/json"
	"kubajaru/rest-api-example/internal/model"
	"kubajaru/rest-api-example/internal/service"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type TaskController struct {
	service *service.TaskService
}

func NewTaskController(s *service.TaskService) *TaskController {
	return &TaskController{service: s}
}

func (c *TaskController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", c.handleTasks)
	mux.HandleFunc("/tasks/", c.handleTaskByID)
}

func (c *TaskController) handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks := c.service.GetAll()
		writeJSON(w, http.StatusOK, tasks)
	case http.MethodPost:
		var task model.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		created := c.service.Create(task)
		slog.Info("Task created", "id", created.ID)
		writeJSON(w, http.StatusCreated, created)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *TaskController) handleTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		task, ok := c.service.GetByID(id)
		if !ok {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		writeJSON(w, http.StatusOK, task)
	case http.MethodPut:
		var task model.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		updated, ok := c.service.Update(id, task)
		if !ok {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		slog.Info("Task updated", "id", id)
		writeJSON(w, http.StatusOK, updated)
	case http.MethodDelete:
		if !c.service.Delete(id) {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		slog.Info("Task deleted", "id", id)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
