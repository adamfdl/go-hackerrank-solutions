package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	tc := strings.Split("1 1 3 1 2 1 3 3 3 3", " ")

	realtc := []int32{}
	for _, numstr := range tc {
		num, _ := strconv.Atoi(numstr)
		realtc = append(realtc, int32(num))
	}

	pairs := sockMerchant(9, realtc)
	fmt.Println(pairs)
}

func sockMerchant(n int32, ar []int32) int32 {

	// Add socks to loop as map
	// Ex. [1,1,1] => [1: 3]
	matchedSocks := map[int32]int32{}
	for i := 0; i < len(ar); i++ {
		if _, ok := matchedSocks[ar[i]]; !ok {
			matchedSocks[ar[i]] = 1
		} else {
			matchedSocks[ar[i]] = matchedSocks[ar[i]] + int32(1)
		}
	}

	fmt.Println(matchedSocks)

	var pairs int32
	for _, val := range matchedSocks {
		// Example: 11/2 equals 5
		pairs = pairs + val/2
	}

	return pairs
}
