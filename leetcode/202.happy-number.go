import (
	"strconv"
)

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 *
 * https://leetcode.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (49.87%)
 * Likes:    1885
 * Dislikes: 400
 * Total Accepted:    476.5K
 * Total Submissions: 955.8K
 * Testcase Example:  '19'
 *
 * Write an algorithm to determine if a number n is "happy".
 *
 * A happy number is a number defined by the following process: Starting with
 * any positive integer, replace the number by the sum of the squares of its
 * digits, and repeat the process until the number equals 1 (where it will
 * stay), or it loops endlessly in a cycle which does not include 1. Those
 * numbers for which this process ends in 1 are happy numbers.
 *
 * Return True if n is a happy number, and False if not.
 *
 * Example:
 *
 *
 * Input: 19
 * Output: true
 * Explanation:
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 *
 *
 */

// @lc code=start
func isHappy(n int) bool {
	m := make(map[int]int)

	for n != 1 {
		n = sum(n)
		// if we find a cycle, it will never be a happy number
		if m[n] > 0 {
			return false
		}
		m[n]++
	}

	return n == 1
}

func sum(s int) int {
	str := strconv.Itoa(s)
	sum := 0

	for _, c := range str {
		num, _ := strconv.Atoi(string(c))
		num *= num
		sum += num
	}

	return sum
}

// @lc code=end
