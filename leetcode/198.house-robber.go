/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 *
 * https://leetcode.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (41.69%)
 * Likes:    4169
 * Dislikes: 125
 * Total Accepted:    474.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,1]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed, the only constraint stopping
 * you from robbing each of them is that adjacent houses have security system
 * connected and it will automatically contact the police if two adjacent
 * houses were broken into on the same night.
 *
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 * Example 2:
 *
 *
 * Input: [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house
 * 5 (money = 1).
 * Total amount you can rob = 2 + 9 + 1 = 12.
 *
 *
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	mem := make([]int, n+1)
	mem[0] = 0
	mem[1] = nums[0]

	for i := 1; i < n; i++ {
		// Example 1 with mem:
		// Initial condition: [0, 1, 0, 0, 0]
		// First iteration:   [0, 1, 2, 0, 0]
		// Second iteration:  [0, 1, 2, 4, 0]
		// Third iteration:   [0, 1, 2, 4, 4]
		//
		// Example 2 with mem:
		// Initial condition: [0, 2, 0, 0, 0, 0]
		// First iteration:   [0, 2, 7, 0, 0, 0]
		// Second iteration:  [0, 2, 7, 9, 0, 0]
		// Third iteration:   [0, 2, 7, 9, 12, 0]
		// Fourth iteration:  [0, 2, 7, 9, 12, 12]
		mem[i+1] = max(mem[i-1]+nums[i], mem[i])
	}

	return mem[n]
}

func max(nums ...int) int {
	_max := nums[0]
	for _, v := range nums[1:] {
		if v > _max {
			_max = v
		}
	}

	return _max
}

// @lc code=end
