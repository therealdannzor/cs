/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 *
 * https://leetcode.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (45.59%)
 * Likes:    3505
 * Dislikes: 83
 * Total Accepted:    585.1K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * Given a binary tree, check whether it is a mirror of itself (ie, symmetric
 * around its center).
 *
 * For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
 *
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 *
 *
 *
 *
 * But the following [1,2,2,null,3,null,3] is not:
 *
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 *
 *
 *
 * Follow up: Solve it both recursively and iteratively.
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
func isSymmetric(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil || root.Right == nil {
		return false
	}

	left := traverseLeft(root.Left, false)
	right := traverseLeft(root.Right, true)

	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

func traverseLeft(node *TreeNode, invert bool) []int {
	result := make([]int, 0)

	if node == nil {
		result = append(result, 0)
		return result
	}

	result = append(result, node.Val)

	if invert {
		right := traverseLeft(node.Right, invert)
		result = append(result, right...)
		left := traverseLeft(node.Left, invert)
		result = append(result, left...)
		return result
	}

	left := traverseLeft(node.Left, invert)
	result = append(result, left...)
	right := traverseLeft(node.Right, invert)
	result = append(result, right...)
	return result
}

// @lc code=end
