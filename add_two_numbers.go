package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	addTwoNumbers(&l1, &l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	val := decimal(l1) + decimal(l2)
	if val == 0 {
		return &ListNode{val, nil}
	}
	result := new(ListNode)
	for val > 0 {
		push(result, val%10)
		val /= 10
	}

	return result.Next
}

func decimal(l *ListNode) int {
	decimal := 0
	node := l
	for i := 0; node != nil; i++ {
		decimal += node.Val * int(math.Pow(10, float64(i)))
		node = node.Next
	}

	return decimal
}

func push(l *ListNode, v int) {
	node := l

	for node.Next != nil {
		node = node.Next
	}

	node.Next = &ListNode{v, nil}
}
