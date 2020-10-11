package main

import "fmt"

func main() {
	visits := countingValleys(8, "UDDDUDUU")
	fmt.Println(visits)
}

func countingValleys(steps int32, path string) int32 {
	var (
		seaLevel           = 0
		valleyVisits int32 = 0
	)

	for i := 0; i < int(steps); i++ {
		step := string(path[i])

		// Count sealevel according to steps
		if step == "U" {
			seaLevel++
		} else {
			seaLevel--
		}

		// Step is UP AND sea level is 0, then it is a valley visit
		if step == "U" && seaLevel == 0 {
			valleyVisits++
		}
	}

	return valleyVisits
}
