package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (38.06%)
 * Likes:    4158
 * Dislikes: 196
 * Total Accepted:    858.7K
 * Total Submissions: 2.3M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */

// @lc code=start
// Runtime: beats 100% of golang submissions
// Memory: beats 30% of golang submissions
func isValid(s string) bool {
	// simple stack representation
	stack := make([]string, 0)

	if s == "" {
		return true
	}

	if len(s) < 2 {
		return false
	}

	for i, v := range s {
		// find all opening parantheses
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, string(v))

		} else {
			// short circuit if we are missing an opening paranthesis
			if len(stack) == 0 {
				return false
			}

			switch string(v) {
			// check that for every case of closening paranthesis there is a corresponding
			// opening paranthesis at the stop of the stack
			case ")":
				if stack[len(stack)-1] != "(" {
					return false
				}
				// pop stack if we have found a closing paranthesis
				stack = stack[:len(stack)-1]
			case "}":
				if stack[len(stack)-1] != "{" {
					return false
				}
				stack = stack[:len(stack)-1]
			case "]":
				if stack[len(stack)-1] != "[" {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}

	// if there are items left in the stack, we have not closed all parantheses
	return len(stack) == 0
}

// @lc code=end
