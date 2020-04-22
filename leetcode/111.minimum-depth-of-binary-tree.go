/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 *
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (36.57%)
 * Likes:    1166
 * Dislikes: 629
 * Total Accepted:    385K
 * Total Submissions: 1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its minimum depth.
 *
 * The minimum depth is the number of nodes along the shortest path from the
 * root node down to the nearest leaf node.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given binary tree [3,9,20,null,null,15,7],
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * return its minimum depth = 2.
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

/*
 * A corner test case
 * Input: [-9,-3,2,null,4,4,0,-6,null,-5]
 * Result: 3
 *
 *
 *                       -9
 *                      /  \
 *                     /    \
 *                    /      \
 *                   -3       2
 *                  / \       / \
 *                 /   \     /   \
 *                /     \   /     \
 *              nil      4  4      0
 *                      / \       / \
 *                     /   \     /   \
 *                    /     \   /     \
 *                   -6    nil  -5     nil
 *
 *
 */

// This problem is not "just" the opposite of the `maxDepth` problem encountered earlier.
// In that case, we can simply crawl the levels and add one for every (existing) level we encounter.
// Should we do that here, then the solution to the corner case above would - erroneously - be 2.
// This is because the left child of -3 is empty and we wouldn't care what happens to the right
// child (and its children).
//
// We need a way to distinguish between actual leaf nodes and non-nil nodes. A proposed algorithm
// is the following:
//
//     1. If an existing node has no children its height is 1 since it is non-nil by definition
//     2. If a node does not have a left child, we will only check the depth of its right child;
//        the same conversely for a right child
//     3. Compare between two lineages and return the shorter one plus 1 (include the parent itself)
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// no children (1)
	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftBranch := minDepth(root.Left)
	rightBranch := minDepth(root.Right)

	// no left side offspring (2)
	if leftBranch == 0 {
		return rightBranch + 1
	}
	// no right side offspring (2)
	if rightBranch == 0 {
		return leftBranch + 1
	}

	// find the shortest lineage (3)
	return min(leftBranch, rightBranch) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
