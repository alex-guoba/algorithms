package bt

import (
	"fmt"
	"testing"
)

func TestRightSideView(t *testing.T) {
	// t.Fatal("not implemented")

	node3 := &TreeNode{3, nil, nil}
	node9 := &TreeNode{9, nil, nil}
	node20 := &TreeNode{20, nil, nil}
	node15 := &TreeNode{15, nil, nil}
	node7 := &TreeNode{7, nil, nil}

	node3.Left = node9
	node3.Right = node20

	node20.Left = node15
	node20.Right = node7

	result := postorderTraversalRecur(node3)

	fmt.Println(result)

}
