package file

import (
	"bufio"
	"fmt"
	"os"
)

const fileName = "file/numbers.txt"

func ProcessFile() (lines []string, err error) {
	file, err := openFile()
	if err != nil {
		return nil, err
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			err = fmt.Errorf("error closing file '%s': %w", fileName, err)
		}
	}()

	lines, err = readFile(file)
	return
}

func openFile() (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %w", fileName, err)
	}
	return file, nil
}

func readFile(file *os.File) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file '%s': %w", fileName, err)
	}

	return lines, nil
}
