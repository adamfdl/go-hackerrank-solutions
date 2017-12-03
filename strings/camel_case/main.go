package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)

	result := countCamelCase(input)
	fmt.Printf("%d", result)
}

// Function returns the number of camel cases
// +1 for the first word
func countCamelCase(input string) int {
	var count int
	for _, value := range input {
		if unicode.IsUpper(value) {
			count++
		}
	}
	return count + 1
}
