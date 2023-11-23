package mlist

import (
	"testing"
)

func TestMerge(t *testing.T) {
	l := MakeList([]int{1, 2, 4})
	r := MakeList([]int{1, 3, 4})

	h := mergeTwoLists(l, r)
	if !CompareList(h, MakeList([]int{1, 1, 2, 3, 4, 4})) {
		t.Errorf("error")
	}
}

func TestMergeI(t *testing.T) {
	l := MakeList([]int{1, 1, 1})
	r := MakeList([]int{2, 3, 4})

	h := mergeTwoLists(l, r)
	if !CompareList(h, MakeList([]int{1, 1, 1, 2, 3, 4})) {
		t.Errorf("error")
	}
}
