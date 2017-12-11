package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)

	result := countCosmicRadiated(input)
	fmt.Printf("%d", result)
}

func countCosmicRadiated(message string) int {
	var radiated int

	var counter int
	for i := 0; i < len(message); i++ {
		if counter == 0 {
			if !strings.EqualFold(string(message[i]), "S") {
				radiated++
			}
		} else if counter == 1 {
			if !strings.EqualFold(string(message[i]), "O") {
				radiated++
			}
		} else if counter == 2 {
			if !strings.EqualFold(string(message[i]), "S") {
				radiated++
			}
		}

		counter++

		if counter == 3 {
			counter = 0
		}
	}

	return radiated
}
