package dsp

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisibleSumPairs(t *testing.T) {
	tc := strings.Split("36 46 25 97 57 14 21 50 75 58 54 32 73 11 36 22 95 46 54 61", " ")

	realtc := []int32{}
	for _, numstr := range tc {
		num, _ := strconv.Atoi(numstr)
		realtc = append(realtc, int32(num))
	}

	result := divisibleSumPairs(20, 7, realtc)
	assert.Equal(t, int32(16), result)
}

func TestDivisibleSumPairs_2(t *testing.T) {
	tc := strings.Split("1 3 2 6 1 2", " ")

	realtc := []int32{}
	for _, numstr := range tc {
		num, _ := strconv.Atoi(numstr)
		realtc = append(realtc, int32(num))
	}

	result := divisibleSumPairs(6, 3, realtc)
	assert.Equal(t, int32(5), result)
}
