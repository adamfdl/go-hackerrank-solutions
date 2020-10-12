package main

import "fmt"

func main() {
	result := repeatedString("a", 1000000000000)
	fmt.Printf("Result: %v\n", result)
}

// (Basically) Occurce of A in s * n
// https://www.youtube.com/watch?v=d5xLkTvJDWY&ab_channel=JAVAAID-CodingInterviewPreparation
/*
	Formula:
	(n/len(s)) + frequency of a in s + freq of in (n % len(s))
 **/

func repeatedString(s string, n int64) int64 {

	var freqOfAsInS int64
	for i := range s {
		char := string(s[i])
		if char == "a" {
			freqOfAsInS++
		}
	}

	// how many times s should be repeated
	divided := n / int64(len(s))
	reminder := n % int64(len(s))

	var freqOfAsInReminder int64
	var i int64
	for i = 0; i < reminder; i++ {
		char := string(s[i])
		if char == "a" {
			freqOfAsInReminder++
		}
	}

	// fmt.Println(divided)
	// fmt.Println(freqOfAsInS)
	// fmt.Println(freqOfAsInReminder)

	return divided*freqOfAsInS + freqOfAsInReminder
}

func freqOfAs(s string) int64 {
	var (
		aOccurence int64
	)

	for index := range s {
		char := string(s[index])
		if char == "a" {
			aOccurence++
		}
	}

	return aOccurence
}

// Brute force will not work
func repeatedStringBruteForce(s string, n int64) int64 {
	if int64(len(s)) >= n {
		return freqOfAs(s)
	}

	for int64(len(s)) < n {
		s += s
	}

	return freqOfAs(s)
}
