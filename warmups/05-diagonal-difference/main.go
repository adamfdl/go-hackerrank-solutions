package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)

	matrix := initMatrix(input)

	var diagA, diagB int
	for i := 0; i < len(matrix); i++ {
		diagA = diagA + matrix[i][i]
		diagB = diagB + matrix[i][len(matrix)-i-1]
	}

	result := diagA - diagB
	fmt.Printf("%v", math.Abs(float64(result)))
}

func initMatrix(size int) [][]int {
	scanner := bufio.NewScanner(os.Stdin)

	// Init multidimensional array
	values := [][]int{}

	for i := 0; i < size; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			tokens := strings.Split(line, " ")

			// Init matrix rows
			arrayNum := []int{}
			for k := range tokens {
				converts, _ := strconv.Atoi(tokens[k])

				// Append ints to rows
				arrayNum = append(arrayNum, converts)
			}

			// Append rows to multidimendisonal array
			values = append(values, arrayNum)
		}
	}

	return values
}
