/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 *
 * https://leetcode.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (39.83%)
 * Likes:    1576
 * Dislikes: 452
 * Total Accepted:    433.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * Given a binary tree and a sum, determine if the tree has a root-to-leaf path
 * such that adding up all the values along the path equals the given sum.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given the below binary tree and sum = 22,
 *
 *
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \      \
 * 7    2      1
 *
 *
 * return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	remainder := sum - root.Val

	if isLeaf(root) && remainder == 0 {
		return true
	}

	return hasPathSum(root.Left, remainder) || hasPathSum(root.Right, remainder)
}

func isLeaf(root *TreeNode) bool {
	if root != nil && root.Left == nil && root.Right == nil {
		return true
	}

	return false
}

// @lc code=end
