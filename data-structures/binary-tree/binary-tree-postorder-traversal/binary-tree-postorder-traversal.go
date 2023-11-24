/*
题目地址(145. 二叉树的后序遍历)

https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

## 题目描述

```
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

```

*/

package bt

// 使用递归实现后序遍历
func postorderTraversalRecur(root *TreeNode) []int {
	result := []int{}

	result = postOrderRecur(root, result)

	return result
}

func postOrderRecur(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}
	if node.Left != nil {
		result = postOrderRecur(node.Left, result)
	}

	if node.Right != nil {
		result = postOrderRecur(node.Right, result)
	}
	result = append(result, node.Val)
	return result
}

// 使用遍历方式实现
// 我们也可以用迭代的方式实现方法一的递归函数，两种方式是等价的，区别在于递归的时候隐式地维护了一个栈，
// 而我们在迭代的时候需要显式地将这个栈模拟出来，其余的实现与细节都相同.
// TODO: 未完成
func postorderTraversalIter(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}

	if root == nil {
		return result
	}

	stack = append(stack, root)
	for len(stack) > 0 {
		if stack.Left != nil {
		}
	}

	return result
}
