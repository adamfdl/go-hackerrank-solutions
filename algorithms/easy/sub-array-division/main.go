package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	tc := strings.Split("2 5 1 3 4 4 3 5 1 1 2 1 4 1 3 3 4 2 1", " ")

	fmt.Println(tc)

	realtc := []int32{}
	for _, numstr := range tc {
		num, _ := strconv.Atoi(numstr)
		realtc = append(realtc, int32(num))
	}

	data := birthday(realtc, 18, 7)
	fmt.Println(data)
}

func birthday(s []int32, d int32, m int32) int32 {
	var counter int32

	if len(s) == 1 && s[0] == d {
		return 1
	}

	for i := 0; i < len(s); i++ {
		// Check bounds
		if i+int(m) > len(s) {
			break
		}

		var tempD int32
		// Count
		for j := 0; j < int(m); j++ {
			tempD = tempD + s[i+j]
		}

		if tempD == d {
			counter++
		}
	}

	return counter
}
