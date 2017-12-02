package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	aliceInput := inputToArray()
	bobInput := inputToArray()

	aliceScore := convertToInt(aliceInput)
	bobScore := convertToInt(bobInput)

	var aliceFinalScore int
	var bobFinalScore int

	for i := 0; i < 3; i++ {
		if aliceScore[i] > bobScore[i] {
			aliceFinalScore++
		} else if aliceScore[i] < bobScore[i] {
			bobFinalScore++
		}
	}

	fmt.Printf("%v %v", aliceFinalScore, bobFinalScore)
}

func inputToArray() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var arrayString []string
	if scanner.Scan() {
		line := scanner.Text()
		arrayString = strings.Split(line, " ")
	}
	return arrayString
}

func convertToInt(score []string) []int {
	var arrayNum []int
	for i := 0; i < len(score); i++ {
		converts, _ := strconv.Atoi(score[i])
		arrayNum = append(arrayNum, converts)
	}
	return arrayNum
}
