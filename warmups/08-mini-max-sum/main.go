package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arrayNumbers := reader()
	miniMaxSum := MiniMaxSum(arrayNumbers)

	min, max := MinMax(miniMaxSum)
	fmt.Printf("%d %d", min, max)
}

func reader() []int {
	scanner := bufio.NewScanner(os.Stdin)

	var arrayString []string
	if scanner.Scan() {
		line := scanner.Text()
		arrayString = strings.Split(line, " ")
	}

	var arrayNum []int
	for i := 0; i < 5; i++ {
		converts, _ := strconv.Atoi(arrayString[i])
		arrayNum = append(arrayNum, converts)
	}

	return arrayNum
}

func MiniMaxSum(arrayNum []int) []int {
	var sum int
	var result []int
	for i := 0; i < 5; i++ {
		sum = 0
		for j := 0; j < 5; j++ {
			sum = sum + arrayNum[j]
		}
		sum = sum - arrayNum[i]
		result = append(result, sum)
	}
	return result
}

func MinMax(arrayNum []int) (int, int) {
	var min int
	var max int

	max = arrayNum[1]
	for _, value := range arrayNum {
		if value > max {
			max = value
		}
	}

	min = arrayNum[1]
	for _, value := range arrayNum {
		if value < min {
			min = value
		}
	}

	return min, max
}
