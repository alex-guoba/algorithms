/*
题目地址(102. 二叉树的层序遍历)

https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

## 题目描述

```
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

```
*/

package bt

// 使用BFS遍历即可。重点在于如何区分level数
func levelOrder(root *TreeNode) [][]int {
	var result = [][]int{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	level := 0
	for {
		if len(queue) == 0 {
			break
		}

		lastLevelNum := len(queue)
		nodes := []int{}
		for i := 0; i < lastLevelNum; i++ {
			t := queue[i]

			nodes = append(nodes, t.Val)
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
		}
		// discard all last level nodes
		queue = queue[lastLevelNum:]

		// next level
		result = append(result, nodes)
		level++
	}

	return result
}

// TODO: 使用DFS也可以实现，但需要记住level
