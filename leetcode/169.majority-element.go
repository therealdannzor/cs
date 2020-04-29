/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (56.47%)
 * Likes:    2720
 * Dislikes: 210
 * Total Accepted:    542.7K
 * Total Submissions: 960.9K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */

// @lc code=start
// An alternative to bruteforce with hashmap,
// use Boyer-Moore majority vote algorithm:
//
// Initialize an element m and a counter i with i = 0
// For each element x of the input sequence:
//   If i = 0, then assign m = x and i = 1
//   else if m = x, then assign i = i + 1
//   else assign i = i - 1
// Return m
func majorityElement(nums []int) int {
	var m int
	var i int

	for _, x := range nums {
		if i == 0 {
			m = x
			i++
			continue
		}
		if m == x {
			i++
		} else {
			i--
		}
	}

	return m
}

// @lc code=end
