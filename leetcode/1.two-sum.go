package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (45.09%)
 * Likes:    13585
 * Dislikes: 494
 * Total Accepted:    2.6M
 * Total Submissions: 5.7M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// hist contains references to the complement of each number
	hist := make(map[int]int)

	for index, value := range nums {
		compl := target - value
		if storedIndex, ok := hist[compl]; ok {
			return []int{storedIndex, index}
		}
		hist[value] = index
	}

	return nil
}

// @lc code=end

func run() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	expectedRes := []int{0, 1}

	if !testEq(res, expectedRes) {
		fmt.Printf("Test FAILED. Expected: %v, got: %v\n", expectedRes, res)
		return
	}
	fmt.Println("Test passed.")
	return
}

func testEq(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
