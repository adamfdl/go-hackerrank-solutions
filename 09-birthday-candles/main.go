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

	candles := reader(input)
	blowAbleCandles := CountBlowableCandles(candles)

	fmt.Printf("%d", blowAbleCandles)
}

func reader(candles int) []int {
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()

	arrayString := strings.Split(string(bytes), " ")

	var arrayNum []int
	for i := 0; i < candles; i++ {
		converts, _ := strconv.Atoi(arrayString[i])
		arrayNum = append(arrayNum, converts)
	}

	return arrayNum
}

func CountBlowableCandles(candles []int) int {
	max := candles[0]
	blowable := 0
	for _, height := range candles {
		if height > max {
			max = height
			blowable = 1
		} else if height == max {
			blowable++
		}
	}
	return blowable
}
