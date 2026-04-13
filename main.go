package main

import (
	"fmt"
	"os"
)

type Command string

func checkNumberOfArgs(args []string, expected int) error {
	if len(args) != expected {
		return fmt.Errorf("expected %d arguments, got %d", expected, len(args))
	}
	return nil
}

func main() {

	const (
		CommandAdd Command = "add"
		CommandUpdate Command = "update"
		CommandDelete Command = "delete"
		CommandMarkInProgress Command = "mark-in-progress"
		CommandMarkDone Command = "mark-done"
		CommandList Command = "list"
		CommandHelp Command = "help"
	)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Wrong number of arguments")
		os.Exit(1)
	}

	args := os.Args[2:]

	commands := map[Command]func() error{
		CommandAdd: func() error {
			if err := checkNumberOfArgs(args, 1); err != nil {
				return err
			}
			return addTask(args[0])
		},
		CommandUpdate: func() error {
			if err := checkNumberOfArgs(args, 2); err != nil {
				return err
			}
			return updateTask(args[0], args[1])
		},
		CommandDelete: func() error {
			if err := checkNumberOfArgs(args, 1); err != nil {
				return err
			}
			return deleteTask(args[0])
		},
		CommandMarkInProgress: func() error {
			if err := checkNumberOfArgs(args, 1); err != nil {
				return err
			}
			return markTaskInProgress(args[0])
		},
		CommandMarkDone: func() error {
			if err := checkNumberOfArgs(args, 1); err != nil {
				return err
			}
			return markTaskDone(args[0])
		},
		CommandList: func() error {
			if err := checkNumberOfArgs(args, 1); err != nil {
				return listTasks("")
			}
			return listTasks(TaskStatus(args[0]))
		},
		CommandHelp: func() error {
			fmt.Println("Available commands:")
			fmt.Println("--------------------------------")
			fmt.Println("add <description> - Add a new task")
			fmt.Println("update <id> <description> - Update a task")
			fmt.Println("delete <id> - Delete a task")
			fmt.Println("mark-in-progress <id> - Mark a task as in progress")
			fmt.Println("mark-done <id> - Mark a task as done")
			fmt.Println("list - List all tasks")
			fmt.Println("list <status> - List tasks by status")
			fmt.Println("help - Show this help message")
			fmt.Println("--------------------------------")
			return nil
		},
	}


	
	command := Command(os.Args[1])

	fn, ok := commands[command]

	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	err := fn()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}