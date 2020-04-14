/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 *
 * https://leetcode.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (38.07%)
 * Likes:    1788
 * Dislikes: 3707
 * Total Accepted:    511.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as
 * one sorted array.
 *
 * Note:
 *
 *
 * The number of elements initialized in nums1 and nums2 are m and n
 * respectively.
 * You may assume that nums1 has enough space (size that is greater or equal to
 * m + n) to hold additional elements from nums2.
 *
 *
 * Example:
 *
 *
 * Input:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * Output: [1,2,2,3,5,6]
 *
 */

// @lc code=start
// * Compare the last non-zero elements of both nums1 and num2
// * Insert the greater or equal of the two to the last index of num1
// * We are finished when lastSec is -1 since lastFir is initially in place
func merge(nums1 []int, m int, nums2 []int, n int) {
	// last element of first and second slice, respectively
	lastFir := m - 1
	lastSec := n - 1
	// points to the (sorted) index of nums1
	ptr := len(nums1) - 1

	// as long as we have not exhausted either of the slices with elements
	for lastSec >= 0 && lastFir >= 0 {
		if nums1[lastFir] >= nums2[lastSec] {
			nums1[ptr] = nums1[lastFir]
			lastFir--
		} else {
			nums1[ptr] = nums2[lastSec]
			lastSec--
		}
		// we have narrowed down one element in the to-be sorted nums1, shrink window
		ptr--
	}

	// it seems the starting elements of nums1 are greater than the ones in nums2
	// but since both are sorted we can simply add the elements of nums2 in place of
	// where elements of num1 initially used to be
	for lastSec > -1 {
		nums1[ptr] = nums2[lastSec]
		lastSec--
		ptr--
	}

}

// @lc code=end
