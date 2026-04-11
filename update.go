package main

import "fmt"

// # Updating and deleting tasks
// task-cli update 1 "Buy groceries and cook dinner"
// task-cli delete 1

// # Marking a task as in progress or done
// task-cli mark-in-progress 1
// task-cli mark-done 1

func updateTask(id string) error {
	fmt.Println("Updating task...")
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