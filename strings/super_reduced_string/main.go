package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)

	result := reduceString(input)
	fmt.Println(result)
}

func reduceString(input string) string {
	for i := 1; i < len(input); i++ {
		if string(input[i]) == string(input[i-1]) {
			input = input[:i-1] + input[i+1:]
			i = 0
		}
	}

	if len(input) == 0 {
		return "Empty String"
	}

	return input
}
