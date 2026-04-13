package main

import (
	"errors"
	"fmt"
	"strconv"
)

func deleteTask(id string) error {
	
	tasks, err := loadTasks()

	if err != nil {
		return err
	}

	taskFound := false

	taskID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
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

	fmt.Println("Task deleted successfully")

	return nil
}