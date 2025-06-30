package model

import "fmt"

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (t Task) ToString() string {
	return fmt.Sprintf("Task: %s, isDone: %t", t.Title, t.Done)
}
