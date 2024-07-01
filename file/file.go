package file

import (
	"bufio"
	"fmt"
	"os"
)

const fileName = "file/numbers.txt"

// ProcessFile opens the file, reads its content line by line, and returns the lines.
//
// Returns:
// - A slice of strings containing the lines of the file.
// - An error if there is an issue opening, reading, or closing the file.
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
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return lines, nil
}

// openFile opens the specified file and returns the file handle.
//
// Returns:
// - A pointer to the opened file.
// - An error if there is an issue opening the file.
func openFile() (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %w", fileName, err)
	}

	return file, nil
}

// readFile reads the file line by line and returns a slice of lines.
//
// Returns:
// - A slice of strings containing the lines of the file.
// - An error if there is an issue reading the file.
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
