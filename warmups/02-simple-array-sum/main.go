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

	scanner := bufio.NewScanner(os.Stdin)
	var arrayNumStr []string
	if scanner.Scan() {
		line := scanner.Text()
		arrayNumStr = strings.Split(line, " ")
	}

	var sum int
	for i := 0; i < input; i++ {
		num, _ := strconv.Atoi(arrayNumStr[i])
		sum = sum + num
	}

	fmt.Printf("%v", sum)
}
