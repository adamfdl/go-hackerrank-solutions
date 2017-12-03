package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)

	result := reverseData(readData(input))
	fmt.Printf("%s", result)
}

func readData(input int) []int {
	scanner := bufio.NewScanner(os.Stdin)

	var arrayString []string
	if scanner.Scan() {
		line := scanner.Text()
		arrayString = strings.Split(line, " ")
	}

	var arrayNum []int
	for i := 0; i < input; i++ {
		converts, _ := strconv.Atoi(arrayString[i])
		arrayNum = append(arrayNum, converts)
	}

	return arrayNum
}

func reverseData(array []int) string {
	var reversedData bytes.Buffer
	for i := len(array) - 1; i >= 0; i-- {
		reversedData.WriteString(strconv.Itoa(array[i]) + " ")
	}
	return strings.TrimSpace(reversedData.String())
}
