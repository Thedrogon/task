package task

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func List() error {
	path, err := dataFilePath()
	if err != nil {
		return err
	}

	// If file doesn't exist → empty list
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("No tasks found.")
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 1
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		if len(parts) < 2 {
			continue
		}

		status := "⭕️"
		if parts[1] == "done" {
			status = "✅"
		}

		fmt.Printf("%d. %s %s\n", i, status, parts[0])
		i++
	}

	return scanner.Err()
}

