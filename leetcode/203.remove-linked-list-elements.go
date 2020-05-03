/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 *
 * https://leetcode.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (37.37%)
 * Likes:    1353
 * Dislikes: 80
 * Total Accepted:    303.1K
 * Total Submissions: 810.8K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * Remove all elements from a linked list of integers that have value val.
 *
 * Example:
 *
 *
 * Input:  1->2->6->3->4->5->6, val = 6
 * Output: 1->2->3->4->5
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	for dummy != nil && dummy.Next != nil {
		// edge case: trim leading target values from head
		if dummy.Next.Val == val && dummy.Next == head {
			head = head.Next
			continue
		}

		if dummy.Next.Val == val {
			dummy.Next = dummy.Next.Next
		} else {
			dummy = dummy.Next
		}
	}

	return head
}

// @lc code=end
