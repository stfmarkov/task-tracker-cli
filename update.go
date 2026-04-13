package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// # Updating and deleting tasks
// task-cli update 1 "Buy groceries and cook dinner"
// task-cli delete 1

// # Marking a task as in progress or done
// task-cli mark-in-progress 1
// task-cli mark-done 1

func updateTask(id string, description string) error {
	if description == "" {
		return errors.New("description is required")
	}

	tasks, err := loadTasks()

	if err != nil {
		return err
	}

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	taskFound := false

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			taskFound = true
			break
		}
	}

	if !taskFound {
		return errors.New("task not found")
	}

	err = saveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Println("Task updated successfully")

	return nil
}

func markTaskInProgress(id string) error {
	fmt.Println("Marking task as in progress...")
	return nil
}

func markTaskDone(id string) error {
	fmt.Println("Marking task as done...")
	return nil
}