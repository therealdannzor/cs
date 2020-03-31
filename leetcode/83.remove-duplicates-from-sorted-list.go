/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (44.28%)
 * Likes:    1219
 * Dislikes: 96
 * Total Accepted:    421.8K
 * Total Submissions: 948.6K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted linked list, delete all duplicates such that each element
 * appear only once.
 *
 * Example 1:
 *
 *
 * Input: 1->1->2
 * Output: 1->2
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->2->3->3
 * Output: 1->2->3
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
func deleteDuplicates(head *ListNode) *ListNode {
	registry := make(map[int]int)

	prev := head
	for i := head; i != nil; i = i.Next {
		value := i.Val

		if registry[value] != 0 {
			prev.Next = i.Next
			continue
		}

		registry[value]++
		prev = i
	}

	return head
}

// @lc code=end







