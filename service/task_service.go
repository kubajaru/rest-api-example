package service

import (
	"kubajaru/rest-api-example/model"
	"kubajaru/rest-api-example/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(r *repository.TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

func (s *TaskService) GetAll() []model.Task {
	return s.repo.GetAll()
}

func (s *TaskService) GetByID(id int) (model.Task, bool) {
	return s.repo.GetByID(id)
}

func (s *TaskService) Create(task model.Task) model.Task {
	return s.repo.Create(task)
}

func (s *TaskService) Update(id int, task model.Task) (model.Task, bool) {
	return s.repo.Update(id, task)
}

func (s *TaskService) Delete(id int) bool {
	return s.repo.Delete(id)
}
