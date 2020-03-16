import (
	"math"
)

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (45.64%)
 * Likes:    6477
 * Dislikes: 287
 * Total Accepted:    796.5K
 * Total Submissions: 1.7M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */

// @lc code=start

// Solution rationale:
//
// If the array only contains positive integers, the solution is trivial, and
// we add all numbers up and return the sum of the array.
// Example: [1, 2, 3, 2]: -> 8
//
// If it contains at least one negative integer, we may need to compare whether
// or not it makes sense to include it in the sum or not. It may be the case that
// we have a bridge between two subarrays that makes up a greater sum than if we
// had taken the sum of each of the arrays individually. We sum all of the integers
// across the bridge.
// Example: [1, 2, -2, 1, 2]: -> 4
//
// In contrast, it may be the case that we arrive at a number with a negative value
// that sums up to an integer that is so small that any positive integer would make
// up a greater sum. We choose the largest subarray seen in the array.
// Example: [1, 2, -4, 2, 3]: -> 5
//
// For the base test, we can find a linear solution which stores the biggest sum,
// in this way:
// Given (input):                [-2, 1, -3, 4, -1, 2, 1, -5, 4]
// Algorithm (current best):     [-2, 1, -2, 4,  3, 5, 6,  1, 5]
// Result: 6
func maxSubArray(nums []int) int {
	var currSum int         // zero
	maxSum := math.MinInt32 // smallest possible number to account for negative elements

	for i := 0; i < len(nums); i++ {

		currSum += nums[i]               // add number to current sum
		bestSum := max(nums[i], currSum) // returns the greater of the two, or the second if equal
		// set currSum to the current number if it is greater than the accumulated sum
		if bestSum == nums[i] {
			currSum = nums[i]
		}

		// update the global maximum if necessary
		if bestSum > maxSum {
			maxSum = bestSum
		}
	}

	res := max(currSum, maxSum)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
