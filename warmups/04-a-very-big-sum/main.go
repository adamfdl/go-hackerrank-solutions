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
	fmt.Scanf("%v", &input)

	nums := reader()

	var sum int64
	for i := 0; i < input; i++ {
		converted, _ := strconv.ParseInt(nums[i], 10, 64)
		sum = sum + int64(converted)
	}

	fmt.Printf("%v", sum)
}

func reader() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var arrayString []string
	if scanner.Scan() {
		line := scanner.Text()
		arrayString = strings.Split(line, " ")
	}
	return arrayString
}
