package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type TaskStatus string

const (
	TaskStatusTodo TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in-progress"
	TaskStatusDone TaskStatus = "done"
)

type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	Status TaskStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Task) save() error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	t.ID = maxID + 1

	tasks = append(tasks, *t)

	err = saveTasks(tasks)
	if err != nil {
		return err
	}
	return nil
}

func loadTasks() ([]Task, error) {
	jsonData, err := os.ReadFile("tasks.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(jsonData, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", jsonData, 0600)
	if err != nil {
		return err
	}
	return nil
}