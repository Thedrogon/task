package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/anikchand461/task/internal/task"
)

func printHelp() {
	fmt.Print(`
Task CLI 📝

Usage:
  task <command> [arguments]

Commands:
  add "task name"     Add a new task
  list                List all tasks
  done <id> <id> ...   Mark one or more tasks as done (space-separated)
  clear                Delete all tasks
  help                 Show this help message
  done all             Mark all tasks as done

Examples:
  task add "Buy milk"
  task list
  task done <id> <id>...... Mark one or more tasks as done (space-separated)

  task clear
`)
}

func main() {
	if len(os.Args) < 2 ||
		os.Args[1] == "help" ||
		os.Args[1] == "-h" ||
		os.Args[1] == "--help" {

		printHelp()
		return
	}

	switch os.Args[1] {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task add <task>")
			return
		}
		err := task.Add(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Task added")

	case "list":
		err := task.List()
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "clear":
		err := task.Clear()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("All tasks deleted")

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task done <id | all> [<id> ...]")
			return
		}

		if os.Args[2] == "all" {
			err := task.DoneAll()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("All tasks marked as done")
			return
		}

		// Join all IDs with commas to support: task done 1 3 or task done "1,3"
		idStr := strings.Join(os.Args[2:], ",")
		err := task.Done(idStr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Tasks marked as done")

	default:
		fmt.Println("Unknown command:", os.Args[1])
		printHelp()
	}
}
