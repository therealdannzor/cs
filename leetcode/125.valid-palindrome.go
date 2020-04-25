import (
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (34.07%)
 * Likes:    1007
 * Dislikes: 2680
 * Total Accepted:    529.8K
 * Total Submissions: 1.5M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * Given a string, determine if it is a palindrome, considering only
 * alphanumeric characters and ignoring cases.
 *
 * Note: For the purpose of this problem, we define empty string as valid
 * palindrome.
 *
 * Example 1:
 *
 *
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "race a car"
 * Output: false
 *
 *
 */

// @lc code=start
func isPalindrome(s string) bool {
	alphanums := trimNonAlphanumeric(s)
	reversed := reverse(alphanums)

	return alphanums == reversed
}

func reverse(s string) string {
	length := len(s)
	runes := make([]rune, length)

	for _, ch := range s {
		length--
		runes[length] = ch
	}
	return string(runes)
}

func trimNonAlphanumeric(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))

	for _, ch := range s {
		ch = unicode.ToLower(ch)
		// the rune should either be a letter or a number and not a space,
		// no other symbols are to be written
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) && !unicode.IsSpace(ch) {
			sb.WriteRune(ch)
		}
	}

	return sb.String()
}

// @lc code=end
