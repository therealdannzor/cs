import (
	"strings"
	"unicode/utf8"
)

/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 *
 * https://leetcode.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (30.16%)
 * Likes:    1046
 * Dislikes: 223
 * Total Accepted:    207.8K
 * Total Submissions: 681.9K
 * Testcase Example:  '1'
 *
 * Given a positive integer, return its corresponding column title as appear in
 * an Excel sheet.
 *
 * For example:
 *
 *
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB
 * ⁠   ...
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: "A"
 *
 *
 * Example 2:
 *
 *
 * Input: 28
 * Output: "AB"
 *
 *
 * Example 3:
 *
 *
 * Input: 701
 * Output: "ZY"
 *
 */

// @lc code=start

const alphaLength int = 26

func convertToTitle(n int) string {
	if n < 0 {
		return ""
	}

	// starting at Z instead of A is (was) not a good idea!
	characters := map[int]string{
		0:  "A",
		1:  "B",
		2:  "C",
		3:  "D",
		4:  "E",
		5:  "F",
		6:  "G",
		7:  "H",
		8:  "I",
		9:  "J",
		10: "K",
		11: "L",
		12: "M",
		13: "N",
		14: "O",
		15: "P",
		16: "Q",
		17: "R",
		18: "S",
		19: "T",
		20: "U",
		21: "V",
		22: "W",
		23: "X",
		24: "Y",
		25: "Z",
	}

	var sb strings.Builder
	for n > 0 {
		n--
		charIndex := n % alphaLength
		n /= alphaLength
		sb.WriteString(characters[charIndex])
	}

	return reverse(sb.String())
}

// reverse crunches a string s and returns its reverse
func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, bytes := utf8.DecodeRuneInString(s)
		s = s[bytes:]
		start += bytes
		utf8.EncodeRune(buf[size-start:], r)
	}

	return string(buf)
}

// @lc code=end
