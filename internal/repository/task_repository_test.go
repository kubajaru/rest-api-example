package repository

import (
	"kubajaru/rest-api-example/internal/model"
	"testing"
)

func TestGetAll(t *testing.T) {
	// Start repository
	repo := NewTaskRepository()

	// given
	task := model.Task{
		Title: "sample",
		Done:  false,
	}

	// when
	repo.Create(task)

	// then
	if _, ok := repo.GetByID(1); !ok {
		t.Errorf("Get by id failed")
	}
}
