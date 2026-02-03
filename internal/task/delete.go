package task

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func DeleteTask(taskNumber int) error {
	path, err := dataFilePath()
	if err != nil {
		return err
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if taskNumber < 1 || taskNumber > len(lines) {
		return fmt.Errorf("invalid task number: %d (valid range: 1-%d)", taskNumber, len(lines))
	}

	lines = append(lines[:taskNumber-1], lines[taskNumber:]...)
	
	return os.WriteFile(
		path,
		[]byte(strings.Join(lines, "\n")+"\n"),
		0644,
	)
}