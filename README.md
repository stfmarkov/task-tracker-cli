# task-tracker

A command-line task tracker. Tasks are stored in `tasks.json` in the directory where you run the program. Each task can be `todo`, `in-progress`, or `done`.

## Usage

Run the program with a command as the first argument, for example:

| Command | What it does |
|--------|----------------|
| `add <description>` | Add a new task (starts as `todo`). |
| `update <id> <description>` | Change a task’s text. |
| `delete <id>` | Remove a task. |
| `mark-in-progress <id>` | Set status to `in-progress`. |
| `mark-done <id>` | Set status to `done`. |
| `list` | List all tasks. |
| `list <status>` | List only `todo`, `in-progress`, or `done` tasks. |
| `help` | Show the built-in help. |

Examples:

```bash
task-tracker add "Buy groceries"
task-tracker list
task-tracker list todo
task-tracker mark-in-progress 1
task-tracker mark-done 1
task-tracker update 1 "Buy groceries and cook"
task-tracker delete 1
```
