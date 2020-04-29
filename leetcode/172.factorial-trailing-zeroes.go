/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 *
 * https://leetcode.com/problems/factorial-trailing-zeroes/description/
 *
 * algorithms
 * Easy (37.69%)
 * Likes:    721
 * Dislikes: 998
 * Total Accepted:    196.9K
 * Total Submissions: 522.5K
 * Testcase Example:  '3'
 *
 * Given an integer n, return the number of trailing zeroes in n!.
 *
 * Example 1:
 *
 *
 * Input: 3
 * Output: 0
 * Explanation: 3! = 6, no trailing zero.
 *
 * Example 2:
 *
 *
 * Input: 5
 * Output: 1
 * Explanation: 5! = 120, one trailing zero.
 *
 * Note: Your solution should be in logarithmic time complexity.
 *
 */

// @lc code=start
func trailingZeroes(n int) int {
	zeroes := 0

	// Example: amount of zeros in 125! is
	// 125/5 + 125/25 + 125/125 = 25 + 5 + 1 = 31
	for i := 5; i <= n; i *= 5 {
		zeroes += n / i
	}

	return zeroes
}

// @lc code=end
