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
	)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: task-tracker <command>")
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
			if err := checkNumberOfArgs(args, 1); err != nil {
				return err
			}
			return updateTask(args[0])
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
		CommandList: listTasks,
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