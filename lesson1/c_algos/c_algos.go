package c_algos

import (
	"log"
	"math/rand"
	"time"
)

// IsAutomorph - automorphic function
// 14 exercise
func IsAutomorph(num int) bool {
	if num <= 0 {
		log.Fatalf("Number must be natural, entered number is: %d", num)
	}

	numSquare := num * num
	for num > 0 {
		if (num % 10 != numSquare % 10) {
			return false
		}
		num /= 10; numSquare /= 10
	}
	return true
}


// 13 exercise
func init() {
	rand.Seed(time.Now().UnixNano())
}

func StandartRand(min, max int) int {
	return rand.Intn(max-min) + min
}

func KnuthRand(min, max, a, b int) int {
	x, m := min, max
	// var m = 3 // 2147483647
	if m < 2 {
		log.Fatalf("m var must be >= 2, cur value is: %d", m)
	}
	formulaVars := [...]int{x,a,b}
	for _, e := range formulaVars {
		if e < 0 || e >= m {
			log.Fatalf("vars must be in range of >=0 and < m")
		}
	}
	return (a * x+ b) % m
}


// 12 exercise
func MaxFromInt(nums ...int) int {
	max := nums[0]
	for _, n := range nums {
		if max < n {
			max = n
		}
	}
	return max
}

func MinFromInt(nums ...int) int {
	min := nums[0]
	for _, n := range nums {
		if min > n {
			min = n
		}
	}
	return min
}


// 6 exercise
func YearIdentifier(age int) (int, string) {
	if age > 150 || age <= 0 {
		log.Fatalf("Age must be less than 150 years, u've entered: %d", age)
	}

	switch {
	case age % 100>=11 && age % 100 <=14:
		return age, "лет"
	case age % 10 == 1:
		return age, "год"
	case age % 10 == 2, age % 10 == 3, age % 10 == 4:
		return age, "года"
	default:
		return age, "лет"
	}
}