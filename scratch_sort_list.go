package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var length int
	for p := head; p != nil; p = p.Next {
		length++
	}
	step := 1
	pt := &ListNode{}
	pt.Next = head
	for step < length {
		p, h := pt, pt.Next
		for h != nil {
			h1, l1 := h, step
			for l1 > 0 && h != nil {
				h, l1 = h.Next, l1-1
			}
			if l1 > 0 {
				break
			}
			h2, l2 := h, step
			for l2 > 0 && h != nil {
				h, l2 = h.Next, l2-1
			}
			l1, l2 = step, step-l2
			for l1 > 0 && l2 > 0 {
				if h1.Val < h2.Val {
					p.Next, h1, l1 = h1, h1.Next, l1-1
				} else {
					p.Next, h2, l2 = h2, h2.Next, l2-1
				}
				p = p.Next
			}
			if l1 > 0 {
				p.Next = h1
			} else {
				p.Next = h2
			}
			for l1 > 0 || l2 > 0 {
				p, l1, l2 = p.Next, l1-1, l2-1
			}
			p.Next = h
		}
		step *= 2
	}
	return pt.Next
}

func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)

	head = &ListNode{}
	p := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
	}
	if left != nil {
		p.Next = left
	} else {
		p.Next = right
	}
	return head.Next
}

func makeList(arr []int) *ListNode {
	var pt ListNode
	p := &pt
	for i := range arr {
		p.Next = &ListNode{Val: arr[i]}
		p = p.Next
	}
	return pt.Next
}

func Print(arr []int) {
	l := sortList(makeList(arr))
	for l != nil {
		fmt.Print(l.Val, ",")
		l = l.Next
	}
	fmt.Println()
}

func main() {
	Print([]int{4, 2, 1, 3})
}
