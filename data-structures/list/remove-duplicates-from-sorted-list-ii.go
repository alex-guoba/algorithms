package mlist

/**
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
 * 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
 *
 */

func deleteDuplicates(head *ListNode) *ListNode {
	// 引入头部dumy节点，规避首节点不好处理的问题
	header := &ListNode{-1000, head}
	cur := header

	for cur.Next != nil && cur.Next.Next != nil {
		// 有重复元素
		if cur.Next.Val == cur.Next.Next.Val {
			val := cur.Next.Val
			// delete repeated val
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return header.Next
}
