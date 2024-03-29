/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 *
 * https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
 *
 * algorithms
 * Easy (52.60%)
 * Likes:    1433
 * Dislikes: 562
 * Total Accepted:    372.3K
 * Total Submissions: 702.1K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers that is already sorted in ascending order, find
 * two numbers such that they add up to a specific target number.
 *
 * The function twoSum should return indices of the two numbers such that they
 * add up to the target, where index1 must be less than index2.
 *
 * Note:
 *
 *
 * Your returned answers (both index1 and index2) are not zero-based.
 * You may assume that each input would have exactly one solution and you may
 * not use the same element twice.
 *
 *
 * Example:
 *
 *
 * Input: numbers = [2,7,11,15], target = 9
 * Output: [1,2]
 * Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
 *
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	var result []int

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			result = append(result, left+1, right+1)
			return result
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}

// @lc code=end
