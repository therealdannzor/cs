/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 *
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (46.94%)
 * Likes:    700
 * Dislikes: 194
 * Total Accepted:    264.9K
 * Total Submissions: 555.6K
 * Testcase Example:  '3'
 *
 * Given a non-negative index k where k ≤ 33, return the k^th index row of the
 * Pascal's triangle.
 *
 * Note that the row index starts from 0.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 3
 * Output: [1,3,3,1]
 *
 *
 * Follow up:
 *
 * Could you optimize your algorithm to use only O(k) extra space?
 *
 */

// @lc code=start
// 0: [1]
// 1: [1,1]
// 2: [1,2,1]
// 3: [1,3,3,1]
// 4: [1,4,6,4,1]
func getRow(rowIndex int) []int {
	// define the first two rows (base cases)
	firstRow, secondRow := []int{1}, []int{1, 1}

	switch rowIndex {
	case 0:
		return firstRow
	case 1:
		return secondRow
	}

	// amount of rows to calculate, after the two base-cases
	consecutiveRows := rowIndex - 2
	result := nextRow(secondRow)
	for consecutiveRows > 0 {
		result = nextRow(result)
		consecutiveRows--
	}

	return result
}

func nextRow(currRow []int) []int {
	count := len(currRow)

	// the next row will be one element longer than the previous (current)
	result := make([]int, count+1)
	// add 1's to the first and last element
	last := len(result) - 1
	result[0] = 1
	result[last] = 1

	for i := 1; i < count; i++ {
		result[i] = currRow[i] + currRow[i-1]
	}
	return result
}

// @lc code=end
