package main

import (
	"fmt"
	"os"

	"github.com/anikchand461/task/internal/task"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task <command>")
		fmt.Println("Commands: add, list, done")
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

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task done <id>")
			return
		}
		err := task.Done(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
		}

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

