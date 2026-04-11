package main

import (
	"fmt"
	"os"
)

type Command string

func main() {

	const (
		CommandAdd Command = "add"
		CommandUpdate Command = "update"
		CommandDelete Command = "delete"
		CommandMarkInProgress Command = "mark-in-progress"
		CommandMarkDone Command = "mark-done"
		CommandList Command = "list"
	)

	commands := map[Command]func() error{
		CommandAdd: addTask,
		CommandUpdate: updateTask,
		CommandDelete: deleteTask,
		CommandMarkInProgress: markTaskInProgress,
		CommandMarkDone: markTaskDone,
		CommandList: listTasks,
	}

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: task-tracker <command>")
		os.Exit(1)
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