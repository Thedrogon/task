package task

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Done(idStr string) error {
	// convert id string to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	// get absolute path to ~/.task/tasks.txt
	path, err := dataFilePath()
	if err != nil {
		return err
	}

	// open file
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// read all lines
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// validate id
	if id < 1 || id > len(lines) {
		return nil
	}

	// mark task as done
	parts := strings.Split(lines[id-1], "|")
	if len(parts) >= 2 {
		parts[1] = "done"
		lines[id-1] = strings.Join(parts, "|")
	}

	// write back to file
	return os.WriteFile(
		path,
		[]byte(strings.Join(lines, "\n")+"\n"),
		0644,
	)
}


