package repository

import (
	"kubajaru/rest-api-example/model"
	"sync"
)

type TaskRepository struct {
	mu     sync.Mutex
	tasks  map[int]model.Task
	nextID int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:  make(map[int]model.Task),
		nextID: 1,
	}
}

func (r *TaskRepository) GetAll() []model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks := make([]model.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}

func (r *TaskRepository) GetByID(id int) (model.Task, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.tasks[id]
	return t, ok
}

func (r *TaskRepository) Create(task model.Task) model.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = r.nextID
	r.nextID++
	r.tasks[task.ID] = task
	return task
}

func (r *TaskRepository) Update(id int, task model.Task) (model.Task, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.tasks[id]
	if !ok {
		return model.Task{}, false
	}

	task.ID = id
	r.tasks[id] = task
	return task, true
}

func (r *TaskRepository) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.tasks[id]
	if ok {
		delete(r.tasks, id)
	}
	return ok
}
