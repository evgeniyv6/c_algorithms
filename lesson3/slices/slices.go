package slices

import (
	"math/rand"
	"time"
)


func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerSlice description
// random slice generator
func GenerSlice(num int) []int {
	randNum := make([]int, num)
	for i := 0; i < num; i++ {
		randNum[i] = rand.Intn(9)
	}
	return randNum
}
