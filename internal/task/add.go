package task

import (
	"os"
)

func Add(task string) error {
	path, err := dataFilePath()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(task + "|pending\n")
	return err
}

