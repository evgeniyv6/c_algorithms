package l2

import (
	"fmt"
	"strconv"
)

// CalcImplementerRecur description
func CalcImplementerRecur(x, y int, action, bracket string) int {
	if x > y {
		return 0
	}
	if x ==  y {
		fmt.Printf("OK => Graph route 20 = %s\n", bracket + " x " + action)
		return 1
	}
	return CalcImplementerRecur(x+1,y, action+"+ 1 ", bracket) +
		CalcImplementerRecur(x*2,y,action+") * 2 ", bracket + "(")
}


// CalcImplementerAr description
func CalcImplementerAr(x,y int) int {
	var ar = make([]int, y+1)
	ar[x] = 1  // init one (first) route

	for i:=x; i <= y; i++ {
		ar[i] += ar[i-1]
		if i % 2 == 0 && i / 2 >=x {
			ar[i] +=ar[i/2]
		}
	}
	return ar[y]
}

func divmod(num, denom int) (quo, rem int) {
	return num / denom, num % denom
}

// DecimalToBin description
// 10 -> 2
func DecimalToBin(x int) string {
	var res string
	for x > 0 {
		d,m := divmod(x,2)
		res = strconv.Itoa(m) + res
		x = d
	}
	return res
}


// DecimalToBinRecur description
// 10 -> 2 - рекурсивно
func DecimalToBinRecur(x int) int {
	if x == 0 {
		return 0
	}
	return x % 2 + DecimalToBinRecur(x / 2) * 10
}


// MultiplyRecur description
// Умножение - рекурсивно
func MultiplyRecur(x,y int) int {
	if x < y {
		return MultiplyRecur(y,x)
	}
	if x == 0 {
		return 0
	}
	if x < 0 && y < 0 {
		x=-x;y=-y
	}
	return y + MultiplyRecur(x-1, y)
}

// MultiplyBitOffset description
// Умножение (используя операции побивого И и сдвига вправо влево)
func MultiplyBitOffset(x,y int) int {
	var res int
	var isNegative bool
	if x < 0 && y < 0 {
		x = -x; y = -y
	}
	if x < 0 {
		x = -x
		isNegative = true
	}
	if y < 0 {
		y = -y
		isNegative = true
	}
	for y > 0 {
		if y&1 == 1 {  // нечетное
			res += x
		}
		x = x<<1; y = y>>1
	}
	if isNegative {
		res = - res
	}
	return res
}