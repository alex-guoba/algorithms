package mlist

import (
	"testing"
)

func TestRemove(t *testing.T) {
	// t.Fatal("not implemented")
	header := MakeList([]int{1, 1, 1, 2, 3})

	deleted := deleteDuplicates(header)
	if !CompareList(deleted, MakeList([]int{2, 3})) {
		t.Errorf("error")
	}
}

func TestRemoveI(t *testing.T) {
	// t.Fatal("not implemented")

	header := MakeList([]int{1, 2, 3, 3, 4, 4, 5})

	deleted := deleteDuplicates(header)
	if !CompareList(deleted, MakeList([]int{1, 2, 5})) {
		t.Errorf("error")
	}
}
