package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}

	result := "not pangram"
	if ok := isPangram(input); ok {
		result = "pangram"
	}

	fmt.Printf("%s", result)

}

func isPangram(input string) bool {
	m := make(map[string]int)

	for _, value := range input {
		char := unicode.ToLower(value)
		if string(char) == " " {

		} else if _, ok := m[string(char)]; !ok {
			m[string(char)] = 1
		}
	}

	result := false
	if len(m) == 26 {
		result = true
	}

	return result
}
