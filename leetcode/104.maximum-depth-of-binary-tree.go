/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 *
 * https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (63.97%)
 * Likes:    2145
 * Dislikes: 66
 * Total Accepted:    741.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its maximum depth.
 *
 * The maximum depth is the number of nodes along the longest path from the
 * root node down to the farthest leaf node.
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
 * return its depth = 3.
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

/* Solution rationale
*
*  Consider a subtree t = [3,2,1] of a bigger tree T.
*  The two leaves and its parent look, from a
*  programmatic perspective, like this:
*
*
*                     3
*                    / \
*                   /   \
*                  2     1
*                 / \   / \
*                /   \ /   \
*              nil nil nil nil
*
*		       Levels: 2
*		       Nodes: 3
*
*  A naive approach is to count the nodes and then derive its depth from that. For example,
*  we see that a tree of size 2 to 3 has three levels, 4 to 7 has three levels, and so on.
*  However, this assumes that the tree is filled towards completion. That is, a new leaf is
*  added to complete the tree. This is not a solution and we can see this through a counter example
*  of this with the subtree t = [1,2,null,3,null,4,null,5].
*
*
*                     1
*                    / \
*                   2   nil
*                  / \
*                 3   nil
*                / \
*               4   nil
*              / \
*             5   nil
*            / \
*         nil   nil
*
*         Levels: 5
*         Nodes: 5
*
*
 */

// Iterative approach
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Solution based on breadth-first search
	height := func(node *TreeNode) int {
		// treasureMap is a queue that represents the nodes we need to search through.
		// A search is defined as checking whether a node has one or two children and
		// adding the ones that exists to the treasure map.
		treasureMap := []*TreeNode{node}

		// missions is the amount of times we have gone on a search. Amount of missions
		// corresponds to the levels (or depth) of a tree since at each mission we search
		// through all the nodes that exists based on what we know from a previous level.
		var missions int

		/*
		* We need to explore the tree level by level and this is done by marking all of the
		* nodes of that level on the treasure map. To know when to stop each exploration
		* round, we count the amount of nodes to go through, `amountSearch`.
		 */
		for {
			// contains the amount of nodes we need to search at each level of the tree
			amountSearch := len(treasureMap)

			// we stop iff when there are no nodes to search (all leaves processed)
			if amountSearch == 0 {
				return missions
			}

			// we count every new mission to embark on
			missions++

			// a mission is not done until we have searched through all of the nodes
			// on our map
			for amountSearch > 0 {
				next := treasureMap[0]
				treasureMap = treasureMap[1:]

				if next.Left != nil {
					treasureMap = append(treasureMap, next.Left)
				}
				if next.Right != nil {
					treasureMap = append(treasureMap, next.Right)
				}
				// done with checking a node; one less node to search
				amountSearch--
			}
		}
	}(root)

	return height
}

// @lc code=end
