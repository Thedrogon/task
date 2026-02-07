package task

import (
	"os"
	"strings"
)
func Add(tasks ...string) (int, error) {
	path, err := dataFilePath()
	if err != nil {
		return 0, err
	}

	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	added := 0

	for _, task := range tasks {
		if strings.TrimSpace(task) == "" {
			continue
		}

		_, err := f.WriteString(task + "|pending\n")
		if err != nil {
			return added, err
		}
		added++
	}

	return added, nil
}
