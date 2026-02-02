package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/anikchand461/task/internal/task"
)

func printHelp() {
	fmt.Println(`
Task CLI 📝

Usage:
  task <command> [arguments]

Commands:
  add "task name"     Add a new task
  list                List all tasks
  done <id>            Mark a task as done
  done all             Mark all tasks as done
  edit <id> "new name" Edit a task description
  clear                Delete all tasks
  help                 Show this help message


Examples:
  task add "Buy milk"
  task list
  task done 1
  task done all
  task edit 1 "Buy groceries and milk"
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

		
	case "edit":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task edit <number> \"<new description>\"")
			fmt.Println("Example: task edit 1 \"Buy groceries\"")
			return
		}

		taskNum, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("❌ Error: Invalid task number")
			fmt.Println("Example: task edit 1 \"Updated description\"")
			return
		}

		newDescription := os.Args[3]

		if strings.TrimSpace(newDescription) == "" {
			fmt.Println("❌ Error: Task description cannot be empty")
			return
		}

		if err := task.EditTask(taskNum, newDescription); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			return
		}

		fmt.Printf("✅ Task #%d updated successfully\n", taskNum)

		
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task done <id | all>")
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

		err := task.Done(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Task marked as done")


	default:
		fmt.Println("Unknown command:", os.Args[1])
		printHelp()
	}
}

