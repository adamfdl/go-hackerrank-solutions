package main

import "fmt"

func main() {
	data := breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1})
	fmt.Println(data)
}

func breakingRecords(scores []int32) []int32 {
	var maxScore int32 = scores[0]
	var minScore int32 = scores[0]

	var (
		breakingHighestRecordGames int32
		breakingLowestRecordGame   int32
	)
	for i := 0; i < len(scores); i++ {
		if scores[i] > maxScore {
			maxScore = scores[i]
			breakingHighestRecordGames++
		}

		if scores[i] < minScore {
			minScore = scores[i]
			breakingLowestRecordGame++
		}
	}

	return []int32{breakingHighestRecordGames, breakingLowestRecordGame}
}
