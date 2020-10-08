package main

import "fmt"

func main() {
	countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
}

/*
	s: starting point of sams's house
	t: ending point of sam's house
	a: location of apple tree
	b: location of orange tree
	apples: distances at which each apple falls from the tree
	oranges: distances at which each oranges falls from the tree
*/

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	isBetweenSamsHouse := func(x int32) bool {
		return x >= s && x <= t
	}

	var (
		samsTotalApples  int32
		samsTotalOranges int32
	)

	for i := 0; i < len(apples); i++ {
		dropLocation := a + apples[i]
		if isBetweenSamsHouse(dropLocation) {
			samsTotalApples++
		}
	}

	for i := 0; i < len(oranges); i++ {
		dropLocation := b + oranges[i]
		if isBetweenSamsHouse(dropLocation) {
			samsTotalOranges++
		}
	}

	fmt.Printf("%v\n%v\n", samsTotalApples, samsTotalOranges)
}
