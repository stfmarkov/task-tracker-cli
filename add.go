package main

import (
	"errors"
	"fmt"
	"time"
)

// # Adding a new task
// task-cli add "Buy groceries"
// # Output: Task added successfully (ID: 1)

func addTask(description string) error {
	if description == "" {
		return errors.New("description is required")
	}
	task := Task{
		Description: description,
		Status: TaskStatusTodo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := task.save()
	if err != nil {
		return err
	}
	fmt.Println("Task added successfully (ID: ", task.ID, ")")
	return nil
}