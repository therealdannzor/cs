/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 *
 * https://leetcode.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (50.04%)
 * Likes:    1184
 * Dislikes: 96
 * Total Accepted:    351.6K
 * Total Submissions: 691.3K
 * Testcase Example:  '5'
 *
 * Given a non-negative integer numRows, generate the first numRows of Pascal's
 * triangle.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 5
 * Output:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1],
 * [1,5,10,10,5,1]
 * ]
 *
 *
 */

// @lc code=start
func generate(numRows int) [][]int {
	var result [][]int
	if numRows == 0 {
		return result
	}

	// first row
	singular := []int{1}
	result = append(result, singular)

	var input []int
	// kickoff the row generator with the first row
	input = singular
	for i := 1; i < numRows; i++ {
		curr := nextRow(input)
		result = append(result, curr)
		input = curr
	}

	return result
}

// nextRow returns the next row in the Pascal triangle,
// give the previous one
func nextRow(currRow []int) []int {
	length := len(currRow)

	var row []int

	for i := 0; i < length; i++ {
		// a row starts with a 1
		if i == 0 {
			row = append(row, 1)
			continue
		}

		// sliding window to add the previous row's elements, in index range [1,len-2]
		element := currRow[i] + currRow[i-1]
		row = append(row, element)
	}

	// a row also ends with a 1
	row = append(row, 1)

	return row
}

// @lc code=end
