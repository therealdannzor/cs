/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 *
 * https://leetcode.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (42.52%)
 * Likes:    1927
 * Dislikes: 151
 * Total Accepted:    418.3K
 * Total Submissions: 976.8K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, determine if it is height-balanced.
 *
 * For this problem, a height-balanced binary tree is defined as:
 *
 *
 * a binary tree in which the left and right subtrees of every node differ in
 * height by no more than 1.
 *
 *
 *
 *
 * Example 1:
 *
 * Given the following tree [3,9,20,null,null,15,7]:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * Return true.
 *
 * Example 2:
 *
 * Given the following tree [1,2,2,3,3,null,null,4,4]:
 *
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 *
 *
 * Return false.
 *
 */

// @lc code=start

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	_, res := isBaseBalance(root)
	if res {
		return true
	}

	return false
}

// isBaseBalance checks the base case where a leaf is balanced if
// and only if it is empty or has no children
func isBaseBalance(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	// get the height and whether the two children are balanced
	leftHeight, leftBalanced := isBaseBalance(root.Left)
	rightHeight, rightBalanced := isBaseBalance(root.Right)

	// get the difference of the two heights
	heightGap := heightDifference(leftHeight, rightHeight)

	// the parent has to have two balanced children and with
	// a height difference within the limit, to be considered
	// balanced itself
	if leftBalanced && rightBalanced && heightGap <= 1 {
		// we cound the longest height and add the parent itself
		return max(leftHeight, rightHeight) + 1, true
	}

	return -1, false
}

// heightDifference returns the difference
// of two heights (abs value)
func heightDifference(x, y int) int {
	diff := x - y
	if diff < 0 {
		diff = diff * -1
	}

	return diff
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// @lc code=end
