package main

import (
	"errors"
	"fmt"
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

	err := findTaskAndUpdate(id, TaskUpdate{
		Description: &description,
	})

	if err != nil {
		return err
	}

	fmt.Println("Task updated successfully")

	return nil
}

func markTaskInProgress(id string) error {

	taskStatus := TaskStatusInProgress
	
	err := findTaskAndUpdate(id, TaskUpdate{
		Status: &taskStatus,
	})

	if err != nil {
		return err
	}

	fmt.Println("Task marked as in progress successfully")
	
	return nil
}

func markTaskDone(id string) error {

	taskStatus := TaskStatusDone

	err := findTaskAndUpdate(id, TaskUpdate{
		Status: &taskStatus,
	})

	if err != nil {
		return err
	}

	fmt.Println("Task marked as done successfully")
	return nil
}