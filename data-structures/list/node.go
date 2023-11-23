package mlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// tool function for unit-test
func MakeList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	var header, cur *ListNode
	for _, i := range arr {
		if header == nil {
			header = &ListNode{i, nil}
			cur = header
		} else {
			cur.Next = &ListNode{i, nil}
			cur = cur.Next
		}
	}
	return header
}

func PrintList(h *ListNode) {
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

func CompareList(l *ListNode, r *ListNode) bool {
	for l != nil && r != nil {
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}

	if l != r {
		return false
	}
	return true
}
