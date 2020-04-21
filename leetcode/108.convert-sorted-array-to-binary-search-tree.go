/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (55.07%)
 * Likes:    2006
 * Dislikes: 190
 * Total Accepted:    369.7K
 * Total Submissions: 660K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given an array where elements are sorted in ascending order, convert it to a
 * height balanced BST.
 *
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 *
 * Example:
 *
 *
 * Given the sorted array: [-10,-3,0,5,9],
 *
 * One possible answer is: [0,-3,9,-10,null,5], which represents the following
 * height balanced BST:
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
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

/*
 * Solution is based on the algorithm from GeeksforGeeks: https://www.geeksforgeeks.org/sorted-array-to-balanced-bst/
 *
 * 1) Get the Middle of the array and make it root.
 * 2) Recursively do same for left half and right half.
 *     a) Get the middle of left half and make it left child of the root
 *         created in step 1.
 *     b) Get the middle of right half and make it right child of the
 *         root created in step 1.
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// get the mid element of the slice and make it a root (1)
	midindex, midvalue := find(nums)
	if midindex == -1 {
		return nil
	}
	root := &TreeNode{}
	root.Val = midvalue

	// make the middle of the left half a child of the root (2a)
	leftchild := sortedArrayToBST(nums[:midindex])
	root.Left = leftchild

	// make the middle of the right half a child of the root (2b)
	rightchild := sortedArrayToBST(nums[midindex+1:])
	root.Right = rightchild

	return root
}

// find gets the mid element of a slice and
// returns the index and value of it
func find(nums []int) (int, int) {
	var mid int
	if len(nums) < 1 {
		return -1, -1
	}
	if len(nums)%2 == 0 {
		mid = len(nums)/2 + 1
	}
	mid = len(nums) / 2
	return mid, nums[mid]
}

// @lc code=end
