package stringCalculator

import (
	"fmt"
	"strconv"
	"strings"
)

const commonDelimiter = "_"

// Add calculate sum the numbers as string. In case the negative numbers, show an error message and returned -1.
func Add(numbers string) int {
	if containsCustomDelimiter(numbers) {
		customDelimiter := extractCustomDelimiter(numbers)
		return addWithDelimiters(numbers, "//", commonDelimiter, "\n", commonDelimiter, customDelimiter, commonDelimiter)
	}
	return addWithDelimiters(numbers, ",", commonDelimiter, "\n", commonDelimiter)
}

func containsCustomDelimiter(value string) bool {
	return strings.HasPrefix(value, "//")
}
func extractCustomDelimiter(value string) string {
	splitted := strings.Split(value, "//")
	return splitted[1][:1]
}

func normalizeDelimiters(unnormalized string, newDelimiters ...string) string {
	r := strings.NewReplacer(newDelimiters...)
	result := r.Replace(unnormalized)
	return result
}

func addWithDelimiters(value string, delimiters ...string) int {
	if value == "" {
		return 0
	}
	normalized := normalizeDelimiters(value, delimiters...)
	normalizedSplitted := strings.Split(normalized, commonDelimiter)
	var acc int
	for _, v := range normalizedSplitted {
		i, err := strconv.Atoi(v)
		if i < 0 {
			fmt.Println("negatives not allowed")
			return -1
		}
		if err == nil && i >= 0 {
			acc += i
		}
	}

	return acc
}
