package main

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (51.29%)
 * Likes:    3376
 * Dislikes: 499
 * Total Accepted:    852.9K
 * Total Submissions: 1.7M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */

// @lc code=start
func mergeTwoLists(n1 *ListNode, n2 *ListNode) *ListNode {
	base := &ListNode{}
	tmp := base

	for n1 != nil && n2 != nil {
		if n1.Val <= n2.Val {
			tmp.Next = n1
			n1 = n1.Next
		} else {
			tmp.Next = n2
			n2 = n2.Next
		}
		tmp = tmp.Next
	}

	if n1 != nil {
		tmp.Next = n1
	} else {
		tmp.Next = n2
	}

	return base.Next
}

// @lc code=end
