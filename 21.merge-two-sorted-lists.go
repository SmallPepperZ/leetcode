/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
package main


// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func clean(node *ListNode) *ListNode {
	node.Next = nil
	return node
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	merge(head, list1, list2)
	return head.Next
}

func merge(head *ListNode, list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return head
	}
	if list1 == nil {
		head.Next = list2
		return head
	}
	if list2 == nil {
		head.Next = list1
		return head
	}
	primary := &ListNode{}
	secondary := &ListNode{}
	if list1.Val < list2.Val {
		primary = list1
		secondary = list2
	} else {
		primary = list2
		secondary = list1
	}
	head.Next = primary
	return merge(head.Next, primary.Next, secondary)
}


// @lc code=end
