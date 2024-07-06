package leetcode

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	minPrice, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}
