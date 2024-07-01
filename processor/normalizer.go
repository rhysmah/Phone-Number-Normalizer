package processor

import (
	"fmt"
	"log"
	"regexp"
)

func NormalizeNumbers(rawNumbers []string) ([]string, error) {

	regex, err := regexp.Compile(`\D`)
	if err != nil {
		return nil, fmt.Errorf("error parsing regular expression: %w", err)
	}

	var normalizedNumbers []string

	for _, num := range rawNumbers {
		normalized := regex.ReplaceAllString(num, "")
		normalizedNumbers = append(normalizedNumbers, normalized)
	}

	log.Println("Successfully normalized data")
	return normalizedNumbers, nil
}
