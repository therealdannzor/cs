/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 *
 * https://leetcode.com/problems/single-number/description/
 *
 * algorithms
 * Easy (62.90%)
 * Likes:    4019
 * Dislikes: 150
 * Total Accepted:    783.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[2,2,1]'
 *
 * Given a non-empty array of integers, every element appears twice except for
 * one. Find that single one.
 *
 * Note:
 *
 * Your algorithm should have a linear runtime complexity. Could you implement
 * it without using extra memory?
 *
 * Example 1:
 *
 *
 * Input: [2,2,1]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: [4,1,2,1,2]
 * Output: 4
 *
 *
 */

// @lc code=start
func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	for key := range m {
		// assumption: only one unique element
		if m[key] == 1 {
			return key
		}
	}

	// we should never reach this
	return -1

}
