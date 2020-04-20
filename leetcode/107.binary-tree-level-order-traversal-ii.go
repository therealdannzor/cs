/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (49.91%)
 * Likes:    1117
 * Dislikes: 201
 * Total Accepted:    290.4K
 * Total Submissions: 574K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its bottom-up level order traversal as:
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	reversedResult := crawlSameLevel(root)
	result := reverseSlice(reversedResult)
	return result
}

func crawlSameLevel(root *TreeNode) [][]int {
	var tovisit []*TreeNode
	var res [][]int

	tovisit = append(tovisit, root)
	var amountvisit int
	for {
		amountvisit = len(tovisit)
		if amountvisit == 0 {
			return res
		}

		var samelevelnodes []int
		for _, node := range tovisit {
			samelevelnodes = append(samelevelnodes, node.Val)
		}

		res = append(res, samelevelnodes)

		for amountvisit > 0 {
			next := tovisit[0]
			tovisit = tovisit[1:]

			if next.Left != nil {
				tovisit = append(tovisit, next.Left)
			}
			if next.Right != nil {
				tovisit = append(tovisit, next.Right)
			}
			amountvisit--
		}
	}
}

func reverseSlice(s [][]int) [][]int {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
	return s
}

// @lc code=end
