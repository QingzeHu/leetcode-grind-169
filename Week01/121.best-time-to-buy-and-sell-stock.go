/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	minPrice := int(^uint(0) >> 1)
	maxProfit := 0
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		} else if maxProfit < price-minPrice {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

// @lc code=end

