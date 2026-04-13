package main

import (
	"fmt"
	"time"
)

// # Listing all tasks
// task-cli list

// # Listing tasks by status
// task-cli list done
// task-cli list todo
// task-cli list in-progress

func listTasks(filterByStatus TaskStatus) error {

	fmt.Println("Listing tasks...", filterByStatus)

	tasks, err := loadTasks()

	if err != nil {
		return err
	}

	for _, task := range tasks {

		if filterByStatus != "" && task.Status != filterByStatus {
			continue
		}

		fmt.Printf("ID: %d \n Description: %s \n Status: %s \n CreatedAt: %s \n UpdatedAt: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339), task.UpdatedAt.Format(time.RFC3339))
		fmt.Println("--------------------------------")
	}
	return nil
}