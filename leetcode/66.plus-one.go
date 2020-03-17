/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (41.91%)
 * Likes:    1242
 * Dislikes: 2089
 * Total Accepted:    526.1K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of digits representing a non-negative integer, plus
 * one to the integer.
 *
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contain a single digit.
 *
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 *
 *
 * Example 2:
 *
 *
 * Input: [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 *
 */

// @lc code=start
func plusOne(digits []int) []int {
	last := len(digits) - 1

	// if the last digit is not 9, simply increment that number and bail out
	if digits[last] != 9 {
		digits[len(digits)-1] += 1
		return digits
	}

	// iterate to count the amount of 9s that becomes a zero
	remainder := 0
	for i := last; i > -1; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			remainder += 1
			continue
		}

		// we found something before going through the whole array
		digits[i] += 1
		return digits
	}

	// the whole array contained 9s so now we need to construct a new array
	// which starts with a 1 and with the same amount trailing zeros as there
	// were nines
	res := []int{1}
	for j := 0; j < remainder; j++ {
		res = append(res, 0)
	}

	return res

}

// @lc code=end
