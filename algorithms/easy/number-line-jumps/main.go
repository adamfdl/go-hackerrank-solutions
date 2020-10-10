package main

import "fmt"

func main() {
	res := kangaroo(0, 3, 4, 2)
	fmt.Println(res)
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Starts at the same position and has equal velocity
	if x1 == x2 {
		if v1 == v2 {
			return "YES"
		} else {
			return "NO"
		}
	}

	isKangorooIsInTheSameSpot := func(jumpNumber int) bool {
		return x1+v1*int32(jumpNumber) == x2+v2*int32(jumpNumber)
	}

	result := "NO"
	for i := 0; i < 10000; i++ {
		if isKangorooIsInTheSameSpot(i) {
			return "YES"
		}
	}

	return result
}
