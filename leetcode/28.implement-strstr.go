package main

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (33.73%)
 * Likes:    1293
 * Dislikes: 1698
 * Total Accepted:    582.2K
 * Total Submissions: 1.7M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 *
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 *
 * Example 1:
 *
 *
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 *
 *
 * Clarification:
 *
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 *
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 *
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" || haystack == needle {
		return 0
	}
	if haystack == "" {
		return -1
	}

	// for a haystack of length `H` and needle of length `N`,
	// create a sliding window of length `N` which moves
	// `H-N+1` times over the haystack
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		// slinding window of length N, start from beginning
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// @lc code=end
