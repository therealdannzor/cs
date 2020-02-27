package main

import "fmt"

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 *
 * https://leetcode.com/problems/currRomanSymbol-to-integer/description/
 *
 * algorithms
 * Easy (54.52%)
 * Likes:    1789
 * Dislikes: 3209
 * Total Accepted:    591.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '"III"'
 *
 * Roman numerals are represented by seven different symbols: I, V, X, L, C, D
 * and M.
 *
 *
 * Symbol       Value
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * For example, two is written as II in Roman numeral, just two one's added
 * together. Twelve is written as, XII, which is simply X + II. The number
 * twenty seven is written as XXVII, which is XX + V + II.
 *
 * Roman numerals are usually written largest to smallest from left to right.
 * However, the numeral for four is not IIII. Instead, the number four is
 * written as IV. Because the one is before the five we subtract it making
 * four. The same principle applies to the number nine, which is written as IX.
 * There are six instances where subtraction is used:
 *
 *
 * I can be placed before V (5) and X (10) to make 4 and 9.
 * X can be placed before L (50) and C (100) to make 40 and 90.
 * C can be placed before D (500) and M (1000) to make 400 and 900.
 *
 *
 * Given a currRomanSymbol numeral, convert it to an integer. Input is guaranteed to be
 * within the range from 1 to 3999.
 *
 * Example 1:
 *
 *
 * Input: "III"
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: "IV"
 * Output: 4
 *
 * Example 3:
 *
 *
 * Input: "IX"
 * Output: 9
 *
 * Example 4:
 *
 *
 * Input: "LVIII"
 * Output: 58
 * Explanation: L = 50, V= 5, III = 3.
 *
 *
 * Example 5:
 *
 *
 * Input: "MCMXCIV"
 * Output: 1994
 * Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
 *
 */

// @lc code=start

func romanToInt(s string) int {
	values := translateToInteger(s)

	var res int
	var empty []int
	for i := 0; len(values) > 0; i++ {
		//fmt.Printf("Iteration %v, res: %v\n", i, res)
		//fmt.Println("Values: ", values)
		first := values[0]

		if len(values) == 1 {
			res += values[0]
			values = empty
			continue
		}

		peek := values[1]

		if first >= peek {
			res += first
			values = values[1:]
			continue
		} else {
			res += peek - first
			values = values[2:]
			continue
		}
	}

	fmt.Println("Returning res: ", res)
	return res
}

func translateToInteger(s string) []int {
	var res []int
	var lookup = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	for _, v := range s {
		res = append(res, lookup[string(v)])
	}
	return res
}

// @lc code=end
