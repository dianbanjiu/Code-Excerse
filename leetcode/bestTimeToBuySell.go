package leetcode

func MaxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			diff := prices[j] - prices[i]
			if diff > max {
				max = diff
			}
		}
	}
	return max
}
