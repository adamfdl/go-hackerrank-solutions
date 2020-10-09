package dsp

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var sumPairs int32

	// Iterate the first loop
	for i := 0; i < int(n)-1; i++ {

		// This is to not repeat the already
		// calculated pairs
		j := i + 1

		for j < int(n) {
			if (ar[i]+ar[j])%k == 0 {
				sumPairs++
			}
			j++
		}
	}

	return sumPairs
}
