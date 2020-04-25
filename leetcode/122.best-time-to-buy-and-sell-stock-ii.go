/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Easy (54.33%)
 * Likes:    2084
 * Dislikes: 1684
 * Total Accepted:    555.8K
 * Total Submissions: 997K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * Say you have an array prices for which the i^th element is the price of a
 * given stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete as many
 * transactions as you like (i.e., buy one and sell one share of the stock
 * multiple times).
 *
 * Note: You may not engage in multiple transactions at the same time (i.e.,
 * you must sell the stock before you buy again).
 *
 * Example 1:
 *
 *
 * Input: [7,1,5,3,6,4]
 * Output: 7
 * Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit
 * = 5-1 = 4.
 * Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 =
 * 3.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4,5]
 * Output: 4
 * Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit
 * = 5-1 = 4.
 * Note that you cannot buy on day 1, buy on day 2 and sell them later, as you
 * are
 * engaging multiple transactions at the same time. You must sell before buying
 * again.
 *
 *
 * Example 3:
 *
 *
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 *
 *
 * Constraints:
 *
 *
 * 1 <= prices.length <= 3 * 10 ^ 4
 * 0 <= prices[i] <= 10 ^ 4
 *
 *
 */

// @lc code=start
func maxProfit(prices []int) int {
	var maxProfit int

	for i := 0; i < len(prices)-1; i++ {
		var window [2]int
		window[0] = prices[i]
		window[1] = prices[i+1]
		maxProfit += buyTheDip(window)
	}

	return maxProfit
}

// buyTheDip receives an array of two prices on two contigious days.
// It returns the difference (profit) of buying the dip, i.e. to buy
// on the day with the lower price and selling the day after.
func buyTheDip(prices [2]int) int {
	firstDay := prices[0]
	secondDay := prices[1]

	if secondDay > firstDay {
		return secondDay - firstDay
	}

	return 0
}

// @lc code=end
