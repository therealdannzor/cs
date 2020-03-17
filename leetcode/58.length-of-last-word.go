import (
	"strings"
)

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.45%)
 * Likes:    547
 * Dislikes: 2155
 * Total Accepted:    341.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word (last word means the last
 * appearing word if we loop from left to right) in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a maximal substring consisting of non-space
 * characters only.
 *
 * Example:
 *
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 *
 *
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	// create function which filters for empty spaces
	f := func(c rune) bool {
		return c == ' '
	}

	// split the string based on filter
	words := strings.FieldsFunc(s, f)

	// bail out if it's empty
	if len(words) == 0 {
		return 0
	}

	lastWord := words[len(words)-1]

	return len(lastWord)
}

// @lc code=end
