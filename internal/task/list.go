package task

import (
	"bufio"
	"fmt"
	"os"
)

func List() error {
	path, err := dataFilePath()
	if err != nil {
		return err
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 1
	for scanner.Scan() {
		fmt.Println(i, scanner.Text())
		i++
	}
	return scanner.Err()
}

