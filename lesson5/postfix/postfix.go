package postfix

import (
	"github.com/evgeniyv6/c_algorithms/lesson5/stack"
	"strings"
)

var bracketsPairs = map[rune]rune{'(':')'}
const (
	openBkt Bracket = iota
	closeBkt
	notBkt
)


// Bracket description
type Bracket int

func whichBracket(ch rune) Bracket {
	for k, v := range bracketsPairs {
		switch ch {
		case k:
			return openBkt
		case v:
			return closeBkt
		}
	}
	return notBkt
}


// IsOperator description
func IsOperator(ch rune) bool {
	return strings.ContainsRune("+-*/^()", ch)
}

// GetPriority description
func GetPriority(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	}
	return -1
}


// HasHigherPriority description
func HasHigherPriority(op1 rune, op2 rune) bool {
	op1Prior := GetPriority(op1)
	op2Prior := GetPriority(op2)
	return op1Prior >= op2Prior
}


// exercise4: перевод из инфиксной записи арифметического выражения в постфиксную

// ToPostfix description
func ToPostfix(expr string) string {
	var st stack.Stack
	var postfixExpr string

	for _, ch := range expr {
		if !IsOperator(ch) {
			postfixExpr += string(ch)
		} else if whichBracket(ch) == openBkt {
			st.Push(ch)
		} else if whichBracket(ch) == closeBkt {
			for !st.IsEmpty() && whichBracket(st.Peek().(rune)) != openBkt {
				postfixExpr += string(st.Pop().(rune))
			}
			st.Pop()
		} else {
			for !st.IsEmpty() && whichBracket(st.Peek().(rune)) != openBkt &&
				HasHigherPriority(st.Peek().(rune), ch) {
				postfixExpr += string(st.Pop().(rune))
			}
			st.Push(ch)
		}
	}
	for !st.IsEmpty() {
		postfixExpr += string(st.Pop().(rune))
	}
	return postfixExpr
}