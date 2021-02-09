package linkedlist

import (
	"fmt"
)


// ListNode description
type ListNode struct {
	Val interface{}
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	}
	if l.Next == nil {
		return fmt.Sprintf("%v -> nil", l.Val)
	}
	return fmt.Sprintf("%v -> %s", l.Val, l.Next)
}

// exercise 3: функция, копирующуя односвязный список (то есть создающую в памяти копию односвязного списка без удаления первого списка).

// CopyLinkedListRec description
func CopyLinkedListRec(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newNode := &ListNode{head.Val,nil}
	newNode.Next = CopyLinkedListRec(head.Next)
	return newNode
}

// CopyLinkedList description
func CopyLinkedList(head *ListNode) *ListNode {
	cur := head
	tmp := new(ListNode)
	p := tmp
	for cur != nil {
		p.Next = &ListNode{cur.Val, p.Next}
		p = p.Next
		cur = cur.Next
	}
	return tmp.Next
}
