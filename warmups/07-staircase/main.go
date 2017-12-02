package main

import (
	"bytes"
	"fmt"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)

	result := PrintStairs(input)
	fmt.Printf("%s", result)
}

func PrintStairs(size int) string {
	var buffer bytes.Buffer

	for i := 0; i < size; i++ {

		for k := size - 1; k > i; k-- {
			buffer.WriteString(" ")
		}

		for j := 0; j < i+1; j++ {
			buffer.WriteString("#")
		}

		buffer.WriteString("\n")
	}

	return buffer.String()
}
