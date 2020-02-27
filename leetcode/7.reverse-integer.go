package main

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.63%)
 * Likes:    2882
 * Dislikes: 4514
 * Total Accepted:    958.7K
 * Total Submissions: 3.7M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 *
 */

// @lc code=start
func reverse(x int) int {
	var res int

	if x < 0 {
		asciiString := reverseHelper(strconv.Itoa(-x))
		integer, _ := strconv.Atoi(asciiString)
		res = -integer
	} else {
		asciiString := reverseHelper(strconv.Itoa(x))
		integer, _ := strconv.Atoi(asciiString)
		res = integer
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}

	return res
}

func reverseHelper(s string) (res string) {
	for _, v := range s {
		res = string(v) + res
	}
	return res
}

// @lc code=end
