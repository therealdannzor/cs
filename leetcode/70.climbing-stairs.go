/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 *
 * https://leetcode.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (46.12%)
 * Likes:    3549
 * Dislikes: 116
 * Total Accepted:    592.7K
 * Total Submissions: 1.3M
 * Testcase Example:  '2'
 *
 * You are climbing a stair case. It takes n steps to reach to the top.
 *
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can
 * you climb to the top?
 *
 * Note: Given n will be a positive integer.
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: 2
 * Explanation: There are two ways to climb to the top.
 * 1. 1 step + 1 step
 * 2. 2 steps
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: 3
 * Explanation: There are three ways to climb to the top.
 * 1. 1 step + 1 step + 1 step
 * 2. 1 step + 2 steps
 * 3. 2 steps + 1 step
 *
 *
 */

// @lc code=start
// Inspiration from: https://www.youtube.com/watch?v=5o-kdjv7FD0&list=WL&index=12&t=131s
func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	// use the memoization principle and only store the last two ways in memory
	mem := make([]int, 2)
	mem[0], mem[1] = 1, 1
	for i := 2; i < n+1; i++ {
		last := mem[1]
		mem[1] += mem[0]
		mem[0] = last
	}

	return mem[1]
}

// @lc code=end
