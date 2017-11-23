package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)

	result := readLine(input)

	var negatives, zeroes, positives float64
	for i := 0; i < len(result); i++ {
		if result[i] < 0 {
			negatives++
		} else if result[i] > 0 {
			positives++
		} else if result[i] == 0 {
			zeroes++
		}
	}

	lenResult := float64(len(result))
	fmt.Printf("%f\n%f\n%f", positives/lenResult, negatives/lenResult, zeroes/lenResult)
}

func readLine(length int) []int {
	scanner := bufio.NewScanner(os.Stdin)

	var arrayString []string
	if scanner.Scan() {
		line := scanner.Text()
		arrayString = strings.Split(line, " ")
	}

	var arrayNum []int
	for i := 0; i < length; i++ {
		converts, _ := strconv.Atoi(arrayString[i])
		arrayNum = append(arrayNum, converts)
	}

	return arrayNum
}
