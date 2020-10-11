package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	tc := strings.Split("0 0 0 1 0 0", " ")
	realtc := []int32{}
	for _, numstr := range tc {
		num, _ := strconv.Atoi(numstr)
		realtc = append(realtc, int32(num))
	}

	result := jumpingOnClouds(realtc)
	fmt.Println("---")
	fmt.Println(result)
}

func jumpingOnClouds(c []int32) int32 {
	var jumps int32 = 0
	for i := 0; i < len(c); i++ {

		fmt.Printf("Iteration = %v\n", i)

		// Check index out of bounds and check +2 and
		if i+2 < len(c) && c[i+2] == 0 {
			jumps++
			i = i + 1
			fmt.Printf("New Iteration Jump 2 = %v\n", i+1)
		} else if i+1 < len(c) && c[i+1] == 0 {
			jumps++
			fmt.Printf("New Iteration Jump 1 = %v\n", i)
		}

	}
	return jumps
}
