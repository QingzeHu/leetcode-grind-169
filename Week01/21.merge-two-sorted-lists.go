/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{-1, nil}
	prev := preHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	}
	if list2 != nil {
		prev.Next = list2
	}
	return preHead.Next
}

// @lc code=end

