package main

import (
	"fmt"
	"github.com/evgeniyv6/c_algorithms/lesson5/linkedlist"
	"github.com/evgeniyv6/c_algorithms/lesson5/postfix"
	"github.com/evgeniyv6/c_algorithms/lesson5/queue"
	"github.com/evgeniyv6/c_algorithms/lesson5/stack"
)

// DecimalToBinWithStack description
// exercise 1: перевод из десятичной в двоичную систему счисления с использованием стека
func DecimalToBinWithStack(num int) {
	var st stack.Stack
	var rem int
	for num != 0 {
		rem = num % 2
		num = num / 2
		st.Push(rem)
	}
	for !st.IsEmpty() {
		fmt.Print(st.Pop())
	}
}

// Bracket description
// exercise 2: определение правильной скобочной последовательности
type Bracket int
var bracketsPairs = map[rune]rune{'}':'{', ')':'(', ']':'['}

const (
	openBkt Bracket = iota
	closeBkt
	notBkt
)

func whichBracket(ch rune) Bracket {
	for k, v := range bracketsPairs {
		switch ch {
		case v:
			return openBkt
		case k:
			return closeBkt
		}
	}
	return notBkt
}

func areBracketsBalanced(formula string) bool {
	var st stack.Stack

	for _, ch := range formula {
		switch whichBracket(ch){
		case openBkt:
			st.Push(ch)
		case closeBkt:
			if st.IsEmpty() {
				return false
			}
			if st.Pop() != bracketsPairs[ch] {
				return false
			}
		}
	}
	if !st.IsEmpty() {
		return false
	}
	return true
}

func main() {
	var myStack stack.Stack
	fmt.Println("Stack. Push 1,2,3")
	myStack.Push(1);myStack.Push(2);myStack.Push(3)
	fmt.Println("Stack.Pop")
	myStack.Pop()
	fmt.Println("As a result stack is (lifo): ", myStack.Dump())

	fmt.Println()
	// decimal to binary
	var num = 153214
	fmt.Println("--- EX 1: DECIMAL TO BINARY ---")
	fmt.Printf("Decimal %d = binary ", num)
	DecimalToBinWithStack(num)

	fmt.Println();fmt.Println()
	fmt.Println("--- EX2: BRACKETS ---")
	b1 := "[2/{5*(4+7)}]"; b2 := "{([])}"; b3:= "{})"
	fmt.Printf("Are balanced brackets %s - %v\n", b1,areBracketsBalanced(b1))
	fmt.Printf("Are balanced brackets %s - %v\n", b2,areBracketsBalanced(b2))
	fmt.Printf("Are balanced brackets %s - %v\n", b3,areBracketsBalanced(b3))

	fmt.Println()
	fmt.Println("--- EX 3: LINKED LISTS ---")
	var ll = &linkedlist.ListNode{1, &linkedlist.ListNode{"2", &linkedlist.ListNode{3, nil}}}
	fmt.Println("Original linked list - ",ll)
	fmt.Println("Copied linked list (recursion) - ", linkedlist.CopyLinkedListRec(ll))
	fmt.Println("Original linked list the same - ",ll)
	fmt.Println("Copied linked list (tmp node)",linkedlist.CopyLinkedList(ll))
	fmt.Println("Original linked list the same - ",ll)

	fmt.Println()
	fmt.Println("--- EX 4: CONVERT TO POSTFIX ---")
	expr1 := "a+(b*c-(d/e^f)*g)*h"; expr2 := "(9^7*3+13)/(9*2+2-4)+5+(8-2)"
	fmt.Printf("Infix: %s, Postfix: %s\n", expr1, postfix.ToPostfix(expr1))
	fmt.Printf("Infix: %s, Postfix: %s\n", expr2, postfix.ToPostfix(expr2))

	fmt.Println()
	fmt.Println("--- EX 5: SLICE (array) QUEUE ---")
	var q queue.Queue
	fmt.Println("Queue. Enqueue 1,2,3")
	q.ArrEnqueue(1); q.ArrEnqueue(2); q.ArrEnqueue(3)
	fmt.Println("Queue.Dequeue")
	q.ArrDequeue()
	fmt.Println("As a result queue is (fifo): ", q.Dump())
}
