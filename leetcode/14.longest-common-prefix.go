/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (34.68%)
 * Likes:    2022
 * Dislikes: 1671
 * Total Accepted:    644.9K
 * Total Submissions: 1.9M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	minWord := minString(strs)

	for i, v := range minWord {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(v) {
				return strs[j][:i]
			}
		}
	}

	return minWord
}

func minString(s []string) string {
	if len(s) < 1 {
		return ""
	}

	assumeShortest := s[0]
	for _, word := range s {
		if len(assumeShortest) > len(word) {
			assumeShortest = word
		}
	}

	return assumeShortest
}

// @lc code=end





