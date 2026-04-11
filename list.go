package main

import "fmt"

// # Listing all tasks
// task-cli list

// # Listing tasks by status
// task-cli list done
// task-cli list todo
// task-cli list in-progress

func listTasks() error {
	fmt.Println("Listing tasks...")
	return nil
}