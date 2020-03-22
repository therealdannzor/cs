import (
	"strconv"
)

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (42.22%)
 * Likes:    1462
 * Dislikes: 253
 * Total Accepted:    398.4K
 * Total Submissions: 937.2K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */

// @lc code=start

// * Start by comparing the length of the two strings. Choose the shorter to iterate over.
//
// * In each iteration, we compare both strings' least-significant (right-most) bit and add
//   any carry bit that we may have. We do this until the shorter string has been processed.
//   The result is written to a final string at each iteration.
//
// * Add the rest of the longer string with the carry and expand the result string.
//
// * Finally, expand the string with 1's corresponding to the value of the carry.

func addBinary(a string, b string) string {
	longer, shorter := length(a, b)
	diff := len(longer) - len(shorter)
	var res string
	var carry string
	// start from the last index of the shortest string
	for i := len(shorter) - 1; i > -1; i-- {
		s, c := binarySum(string(longer[i+diff]), string(shorter[i]), carry) // compensate for the longer slice by adding the difference
		carry = c
		res = s + res
	}

	// continue from the index of where the shortest string finished above
	// but now only with the carry and assume values to add are zeros
	for j := diff - 1; j > -1; j-- {
		s, c := binarySum(string(longer[j]), "", carry)
		carry = c
		res = s + res
	}

	if carry == "1" {
		res = "1" + res
	}

	return res
}

func length(a, b string) (string, string) {
	if len(a) > len(b) {
		return a, b
	}
	return b, a
}

// binarySum adds two numbers and a carry and returns the result
func binarySum(i, j, carry string) (string, string) {
	if j == "" {
		j = "0"
	}

	a, _ := strconv.Atoi(i)
	b, _ := strconv.Atoi(j)
	c, _ := strconv.Atoi(carry)

	switch sum := a + b + c; sum {
	case 0:
		return "0", "0"
	case 1:
		return "1", "0"
	case 2:
		return "0", "1"
	case 3:
		return "1", "1"
	}

	return "", ""

}

// @lc code=end
