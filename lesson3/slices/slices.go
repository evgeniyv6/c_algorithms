package slices

import "math/rand"

// GenerSlice description
// random slice generator
func GenerSlice(num int) []int {
	randNum := make([]int, num)
	for i := 0; i < num; i++ {
		randNum[i] = rand.Intn(50)
	}
	return randNum
}
