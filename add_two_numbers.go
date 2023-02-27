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
	val := toDecimal(l1) + toDecimal(l2)
	if val == 0 {
		return &ListNode{val, nil}
	}
	nodes := &ListNode{val % 10, nil}
	val /= 10
	for val > 0 {
		nval := val % 10
		lnode := lastNode(nodes)
		lnode.Next = &ListNode{nval, nil}
		val /= 10
	}

	return nodes
}

func lastNode(l *ListNode) *ListNode {
	var last *ListNode
	for l.Next != nil {
		last = l.Next
		l = l.Next
	}

	if last == nil {
		return l
	}

	return last
}

func toDecimal(l *ListNode) int {
	decimal := 0
	var digits []int
	digits = append(digits, l.Val)
	for l.Next != nil {
		l = l.Next
		digits = append(digits, l.Val)
	}
	for len(digits) > 0 {
		len := len(digits)
		digit := digits[len-1]
		digits = digits[:len-1]
		decimal += digit * int(math.Pow(10, float64(len-1)))
	}
	return decimal
}
